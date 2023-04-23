<script setup>
import { ref } from 'vue';
import { TranslatorFactory } from 'cn-poe-translator';
const factory = TranslatorFactory.Default();
const textTranslator = factory.getTextTranslator();

const input = ref(null);
const output = ref(null);

function translate(){
  const text= input.value;
  if(text){
    output.value = textTranslator.translate(text);
    navigator.clipboard.writeText(output.value);
  }
}
</script>

<template>
  <div class="container">
    <div class="itemTextList">
      <textarea class="itemText" v-model="input"></textarea>
      <textarea class="itemText" v-model="output" disabled></textarea>
    </div>
  </div>
  <button @click="translate">翻译</button>
</template>

<style>
textarea {
  resize: none;
}

.itemTextList>.itemText:nth-child(n+2) {
  margin-left:10px;
}
.itemText {
  height: 500px;
  width: 400px;
}
</style>
