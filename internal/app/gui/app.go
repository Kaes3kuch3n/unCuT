package gui

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
	"os"
	"path"
	"sync"
	"time"
	"uncut/internal/app/uncut/db"
	"uncut/internal/app/uncut/dtos"
	"uncut/internal/app/uncut/entities"
	"uncut/internal/app/uncut/lead"
)

// App struct
type App struct {
	ctx         context.Context
	database    *gorm.DB
	name        string
	appDir      string
	scheduleDir string
}

// NewApp creates a new App application struct
func NewApp(name string) *App {
	return &App{name: name, database: db.Connect("db.sqlite")}
}

// Startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	// Create app directory tree if it doesn't exist
	a.createDirectories()
}

func (a *App) createDirectories() {
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Errorf("unable to access app directory [%w]", err))
	}
	// Create app directory
	a.appDir = path.Join(dir, a.name)
	err = os.MkdirAll(a.appDir, 0755)
	if err != nil {
		panic(fmt.Errorf("unable to create app directory [%w]", err))
	}
	// Create sub-directories
	a.scheduleDir = path.Join(a.appDir, "schedules")
	err = os.MkdirAll(a.scheduleDir, 0755)
	if err != nil {
		panic(fmt.Errorf("unable to create schedule directory [%w]", err))
	}
}

// GetStoredSchedules returns a list of locally stored schedule templates
func (a *App) GetStoredSchedules() []string {
	entries, err := os.ReadDir(a.scheduleDir)
	if err != nil {
		fmt.Printf("error: unable to read schedule directory [%s]\n", err.Error())
		return []string{}
	}

	schedules := make([]string, 0, len(entries))
	for _, entry := range entries {
		content, err := os.ReadFile(path.Join(a.scheduleDir, entry.Name()))
		if err != nil {
			fmt.Printf("error: unable to open schedule file %s [%s]\n", entry.Name(), err.Error())
			return []string{}
		}
		schedules = append(schedules, string(content))
	}

	return schedules
}

func (a *App) GetAdvertisements() []dtos.Ad {
	adMap := db.LoadAds(a.database)

	ads := make([]dtos.Ad, 0, len(adMap))
	for _, ad := range adMap {
		ads = append(ads, dtos.AdToDTO(ad))
	}

	return ads
}

func (a *App) GetScreenTypes() []dtos.ScreenType {
	screenEntities := db.LoadScreenTypes(a.database)

	screens := make([]dtos.ScreenType, 0, len(screenEntities))
	for _, screen := range screenEntities {
		screens = append(screens, dtos.ScreenTypeToDTO(screen))
	}

	return screens
}

func (a *App) GetCinemas(filter string, maxCount int) []dtos.Cinema {
	cinemaEntities := db.GetCinemasByName(a.database, filter, maxCount)

	cinemas := make([]dtos.Cinema, 0, len(cinemaEntities))
	for _, cinema := range cinemaEntities {
		cinemas = append(cinemas, dtos.CinemaToDTO(cinema))
	}

	return cinemas
}

func (a *App) GetDefaultOutputDirectory() (outputDirectory string) {
	outputDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("unable to get user home directory [%w]", err))
	}
	return path.Join(outputDirectory, "Downloads")
}

func (a *App) SelectOutputDirectory() (outputDirectory string) {
	outputDirectory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:     a.GetDefaultOutputDirectory(),
		Title:                "Select output directory",
		CanCreateDirectories: true,
	})
	if err != nil {
		panic(fmt.Errorf("failed to open directory dialog [%w]", err))
	}
	return outputDirectory
}

func (a *App) GetSchedulingStrategies() (strategies []adSchedulingStrategyDescriptor) {
	return adSchedulingStrategies
}

func (a *App) GenerateLeads(
	scheduleJson string,
	contentJson string,
	screeningIds []int,
	outputDirectory string,
	adScheduleStrategyJson string,
) {
	defer a.recoverError()

	ads := db.LoadAds(a.database)

	// Create ad scheduling strategy
	strategy, err := lead.CreateSelectedStrategy([]byte(adScheduleStrategyJson))
	if err != nil {
		panic(fmt.Errorf("failed to create ad scheduling strategy [%w]", err))
	}

	// Create lead template
	template := lead.NewTemplate([]byte(scheduleJson), []byte(contentJson), strategy)

	// Set up output directory
	outputDirectory = path.Join(outputDirectory, time.Now().Format("2006-01-02T15-04-05"))
	err = os.MkdirAll(outputDirectory, 0755)
	if err != nil {
		panic(fmt.Errorf("failed to create output directory [%w]", err))
	}

	// Put lead together from parts
	leads := make(map[string]*lead.Lead, len(screeningIds))
	for _, id := range screeningIds {
		screening := db.LoadScreening(a.database, uint(id))
		l := createLeadForScreening(template, screening, ads, a.database)

		outputPath := path.Join(outputDirectory, fmt.Sprintf("%s.mp4", screening.Date.Format("2006-01-02")))
		leads[outputPath] = l
	}

	var wg sync.WaitGroup
	for out, l := range leads {
		wg.Add(1)
		go func(lead *lead.Lead, outputPath string) {
			defer wg.Done()
			err := lead.Generate(outputPath)
			if err != nil {
				panic(fmt.Errorf("error while generating lead [%w]", err))
			}
		}(l, out)
	}
	wg.Wait()
}

func createLeadForScreening(
	template lead.Template,
	screening *entities.Screening,
	ads entities.AdMap,
	database *gorm.DB,
) (l *lead.Lead) {
	l = lead.New(template)

	// Fill lead with content
	l.AddScreens(screening.Cinema.Screens)

	// TODO: Move hardcoded "upcoming" id screen name to ? (some json file? config?)
	upcomingMovies := db.GetUpcomingMovies(database, screening, template.GetTrailerCount())
	movieScreens := db.GetScreensForMovies(database, screening.Cinema, upcomingMovies)
	l.AddTrailers(upcomingMovies, movieScreens, screening.Cinema.Screens[2])
	l.AddAds(ads)

	return l
}

func (a *App) recoverError() {
	rec := recover()
	if rec == nil {
		return
	}
	err, ok := rec.(error)
	var msg string
	if ok {
		msg = err.Error()
	} else {
		msg = fmt.Sprintf("%v", rec)
	}
	fmt.Printf("Error: %s\n", err)
	runtime.EventsEmit(a.ctx, "error", msg)
}
