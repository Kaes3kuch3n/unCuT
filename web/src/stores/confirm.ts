import { defineStore } from "pinia";

// eslint-disable-next-line @typescript-eslint/no-unused-vars
function noop(_: boolean) {
  return;
}

export const useConfirmStore = defineStore("confirm", {
  state: () => ({
    respond: noop,
    promptText: "",
    show: false,
  }),
  actions: {
    prompt(text: string): Promise<boolean> {
      this.promptText = text;
      this.show = true;
      return new Promise((resolve) => (this.respond = resolve));
    },
    reply(answer: boolean) {
      this.show = false;
      this.promptText = "";
      this.respond(answer);
      this.respond = noop;
    },
  },
});
