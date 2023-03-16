import { ContentType } from "@/types/resource";

const CONTENT_TYPE_PREFIX = "application/vnd.uncut.content-type";
const DATA_TYPE = "application/json";
const DROP_TARGET_CLASS = "drop-target";

interface DragItem {
  contentType: ContentType;
}

export function useDraggable(item: DragItem) {
  function drag(event: DragEvent) {
    if (!event.dataTransfer) return;
    event.dataTransfer.setData(DATA_TYPE, JSON.stringify(item));
    event.dataTransfer.setData(
      `${CONTENT_TYPE_PREFIX}.${item.contentType.toLowerCase()}`,
      ""
    );
  }

  return {
    drag,
  };
}

type ValueOrGetter<T> = T | (() => T);

export function useDropArea(
  allowedContentType: ValueOrGetter<ContentType | null>
) {
  let nestedCounter = 0;

  function dragOver(event: Event) {
    event.preventDefault();
  }

  function dragEnter(event: DragEvent) {
    const dt = event.dataTransfer;
    const allowedType =
      typeof allowedContentType === "function"
        ? allowedContentType()
        : allowedContentType;
    // Check if drag item meets drop requirements
    if (
      !dt?.types.includes(DATA_TYPE) ||
      (allowedType !== null &&
        !dt.types.includes(
          `${CONTENT_TYPE_PREFIX}.${allowedType.toLowerCase()}`
        )) ||
      !event.currentTarget
    )
      return;

    nestedCounter++;
    if (nestedCounter > 1) {
      // Entered sub-area, don't initialize drop area since that was already done when entering the parent
      return;
    }

    // Allow dropping in this area
    event.preventDefault();
    event.currentTarget.addEventListener("dragover", dragOver);
    (event.currentTarget as HTMLElement).classList.add(DROP_TARGET_CLASS);
  }

  function dragLeave(event: DragEvent) {
    if (nestedCounter <= 0) return;
    nestedCounter--;
    if (!event.currentTarget || nestedCounter > 0) return;

    event.currentTarget.removeEventListener("dragover", dragOver);
    (event.currentTarget as HTMLElement).classList.remove(DROP_TARGET_CLASS);
    nestedCounter = 0;
  }

  function drop(event: DragEvent): unknown | null {
    if (!event.dataTransfer) return null;
    // Disable drop area
    dragLeave(event);
    // Retrieve drag data
    const data = event.dataTransfer.getData(DATA_TYPE);
    return JSON.parse(data);
  }

  return {
    dragEnter,
    dragLeave,
    drop,
  };
}
