<script>
  import { Translate } from "../wailsjs/go/main/App.js";
  import { CaptureRegion } from "../wailsjs/go/main/App.js";

  let inputText;
  let resultText = "";
  function translate() {
    Translate(inputText, "id").then((result) => (resultText = result));
  }

  let isDrawing = false;
  let startX = 0, startY = 0;
  let currentX = 0, currentY = 0;

  function onMouseDown(e) {
    isDrawing = true;
    startX = e.clientX;
    startY = e.clientY;
    currentX = startX;
    currentY = startY;
  }

  function onMouseMove(e) {
    if (!isDrawing) {
      return;
    }
    currentX = e.clientX;
    currentY = e.clientY;
  }

  function onMouseUp(e) {
    isDrawing = false;
    CaptureRegion(Math.floor(boxLeft), Math.floor(boxTop), Math.floor(boxWidth), Math.floor(boxHeight));
  }

  $: boxLeft = Math.min(startX, currentX);
  $: boxTop = Math.min(startY, currentY);
  $: boxWidth = Math.abs(currentX - startX);
  $: boxHeight = Math.abs(currentY - startY);
</script>

<main>
  <div
    style="position:fixed; top:0; left:0; width:100vw; height:100vh; background-color: rgba(0,0,0,0.01);"
    on:mousedown={onMouseDown}
    on:mousemove={onMouseMove}
    on:mouseup={onMouseUp}
  >
    {#if isDrawing}
      <div
        style="
        position:fixed;
        left:{boxLeft}px;
        top:{boxTop}px;
        width:{boxWidth}px;
        height:{boxHeight}px;
        border: 2px solid white;
        pointer-events:none;
      "
      ></div>
    {/if}
  </div>
  
  <div class="result" id="result" placeholder="">{resultText}</div>

  <div class="input-box" id="input">
    <input
      autocomplete="off"
      bind:value={inputText}
      class="input"
      id="inputText"
      type="text"
    />
    <button class="btn" on:click={translate}>Translate</button>
  </div>
</main>

<style>
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 100px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>