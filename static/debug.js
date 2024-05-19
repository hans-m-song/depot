// https://htmx.org/api/#logger
htmx.logger = (_elt, event, data) => {
  const whitelist = ["htmx:load", "htmx:trigger", "htmx:confirm"];

  if (
    event.startsWith("app:") ||
    event.startsWith("htmx:xhr:") ||
    event.includes("error") ||
    whitelist.includes(event)
  ) {
    console.log(`[${event}]`, data);
  }
};
