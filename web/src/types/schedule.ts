import { ScheduleSlot } from "./scheduleSlot";

export type Schedule = ScheduleSlot[];

export interface ScheduleTemplate {
  name: string;
  scheduleTemplate: Schedule;
}
