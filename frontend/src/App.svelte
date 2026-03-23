<script>
  import { Translate } from "../wailsjs/go/main/App.js";

  let inputText;
  let resultText = "";

  function translate() {
    Translate(inputText, "id").then((result) => (resultText = result));
  }
</script>

<main>
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
