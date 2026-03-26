<script>
  import { Translate } from "../wailsjs/go/main/App.js";
  import { CaptureAndTranslate } from "../wailsjs/go/main/App.js";
  import { WindowHide } from "../wailsjs/runtime/runtime.js";

  function onKeyDown(e) {
    if (e.key === "Escape") WindowHide();
  }

  let inputText;
  let resultText = "";
  function translate() {
    Translate(inputText, "id").then((result) => (resultText = result));
  }

  let isDrawing = false;
  let startX, startY;
  function onMouseDown(e) {
    isDrawing = true;
    startX = e.clientX;
    startY = e.clientY;
    width = 0;
    height = 0;
  }

  let width, height;
  let currentX, currentY;

  function onMouseMove(e) {
    if (!isDrawing) return;
    currentX = e.clientX;
    currentY = e.clientY;
    width = Math.abs(currentX - startX);
    height = Math.abs(currentY - startY);
  }

  function onMouseUp(e) {
    isDrawing = false;
    if (width > 5 && height > 5) {
      const captureX = Math.min(startX, currentX);
      const captureY = Math.min(startY, currentY);
      CaptureAndTranslate(captureX, captureY, width, height, "id")
        .then((result) => {
          resultText = result;
        })
        .catch((err) => {
          console.error("Full error:", err);
          resultText = "Error: " + JSON.stringify(err);
        });
    }
  }
</script>

<svelte:window on:keydown={onKeyDown} />
<main>
  <div
    style="position:fixed; top:0; left:0; width:100vw; height:100vh;"
    on:mousedown={onMouseDown}
    on:mousemove={onMouseMove}
    on:mouseup={onMouseUp}
  >
    {#if isDrawing}
      <div
        style="
        position:fixed;
        left:{Math.min(startX, currentX)}px;
        top:{Math.min(startY, currentY)}px;
        width:{width}px;
        height:{height}px;
        border: 2px solid white;
        pointer-events:none;
      "
      ></div>
    {/if}
  </div>
  <div class="result" id="result" placeholder="">{resultText}</div>

</main>

<style>
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }
</style>
