import { createApp } from "vue";
import App from "./App.vue";
import { createPinia } from "pinia";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faTrash } from "@fortawesome/free-solid-svg-icons/faTrash";
import { faClone } from "@fortawesome/free-solid-svg-icons/faClone";
import { faQuestionCircle } from "@fortawesome/free-solid-svg-icons/faQuestionCircle";
import { faCaretRight } from "@fortawesome/free-solid-svg-icons/faCaretRight";
import { faCaretDown } from "@fortawesome/free-solid-svg-icons/faCaretDown";
import { faFilm } from "@fortawesome/free-solid-svg-icons/faFilm";
import { faImage } from "@fortawesome/free-solid-svg-icons/faImage";
import { faChevronDown } from "@fortawesome/free-solid-svg-icons/faChevronDown";
import { faChevronUp } from "@fortawesome/free-solid-svg-icons/faChevronUp";
import { faCheck } from "@fortawesome/free-solid-svg-icons/faCheck";
import { faX } from "@fortawesome/free-solid-svg-icons/faX";
import { faCog } from "@fortawesome/free-solid-svg-icons/faCog";
import { faClapperboard } from "@fortawesome/free-solid-svg-icons/faClapperboard";
import { faFolderOpen } from "@fortawesome/free-solid-svg-icons/faFolderOpen";
import { faCircleXmark } from "@fortawesome/free-solid-svg-icons/faCircleXmark";

library.add(
  faTrash,
  faClone,
  faQuestionCircle,
  faCaretRight,
  faCaretDown,
  faFilm,
  faImage,
  faChevronDown,
  faChevronUp,
  faX,
  faCheck,
  faCog,
  faClapperboard,
  faFolderOpen,
  faCircleXmark
);

const pinia = createPinia();

createApp(App)
  .component("font-awesome-icon", FontAwesomeIcon)
  .use(pinia)
  .mount("#app");
