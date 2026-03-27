<script>
  import { Translate } from "../wailsjs/go/main/App.js";
  import { CaptureAndTranslate } from "../wailsjs/go/main/App.js";
  import { WindowHide } from "../wailsjs/runtime/runtime.js";

  function onKeyDown(e) {
    if (e.key === "Escape") WindowHide();
  }

  let resultText = "";

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
  
  let isLoading = false;
  function onMouseUp(e) {
    isDrawing = false;
    if (width > 5 && height > 5) {
      isLoading = true;
      const captureX = Math.min(startX, currentX);
      const captureY = Math.min(startY, currentY);
      CaptureAndTranslate(captureX, captureY, width, height, "id")
        .then((result) => {
          resultText = result;
          isLoading = false;
        })
        .catch((err) => {
          console.error("Full error:", err);
          resultText = "Error: " + JSON.stringify(err);
          isLoading = false;
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
  <div class="result">
    {#if isLoading}
      <div class="spinner"></div>
    {:else}
      {resultText}
    {/if}
  </div>
</main>

<style>
  .result {
    position: fixed;
    bottom: 20px;
    left: 50%;
    transform: translateX(-50%);
    color: white;
    font-size: 1.2rem;
    background: rgba(0, 0, 0, 0.6);
    padding: 8px 16px;
    border-radius: 8px;
    pointer-events: none;
  }

  .spinner {
    width: 24px;
    height: 24px;
    border: 3px solid rgba(255, 255, 255, 0.3);
    border-top-color: white;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
    margin: 0 auto;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
</style>
