<script setup>
import { ref, computed } from 'vue';
import CnPoeExportDb from "cn-poe-export-db/dist/db.global";

const { accessories, armour, flasks, jewels, weapons, gems } = CnPoeExportDb;
const baseTypesList = [accessories, armour, flasks, jewels, weapons];

const baseTypeMap = new Map();
const baseTypeNameList = [];
const uniqueMap = new Map();
const uniqueNameList = [];

const baseTypeCount = new Map();
for (const baseTypes of baseTypesList) {
  for (const baseType of baseTypes) {
    let zh = baseType.zh;
    if (baseTypeMap.has(zh)) {
      let n = 1;
      if (baseTypeCount.has(zh)) {
        n += baseTypeCount.get(zh);
      }
      baseTypeCount.set(zh, n);
      zh = `${zh}${n}`;
    }
    baseTypeMap.set(zh, baseType.en);
    baseTypeNameList.push(zh);
    if (baseType.uniques) {
      for (const unique of baseType.uniques) {
        let zh = `${unique.zh} ${baseType.zh}`;
        uniqueMap.set(zh, unique.en);
        uniqueNameList.push(zh);
      }
    }
  }
}

const gemMap = new Map();
const gemNameList = [];
for (let gem of gems) {
    gemMap.set(gem.zh, gem.en);
    gemNameList.push(gem.zh);
}

const namesList = [
  baseTypeNameList, uniqueNameList, gemNameList
]

const queryType = ref("0");
const matchedList = computed(() => {
  return namesList[queryType.value];
})

const input = ref(null);
const output = ref(null)
function inputing() {
  answer();
}

function answer(){
  const text = input.value;
  if (queryType.value === "0") {
    if (baseTypeMap.has(text)) {
      output.value = baseTypeMap.get(text);
      return;
    }
  } else if (queryType.value === "1") {
    if (uniqueMap.has(text)) {
      output.value = uniqueMap.get(text);
      return;
    }
  }else if (queryType.value === "2") {
    if (gemMap.has(text)) {
      output.value = gemMap.get(text).replace(" Support","");
      return;
    }
  }
  
  output.value = null;
}

function queryTypeChanged(){
  answer();
}
</script>

<template>
  <div class="container">
    <select id="queryType" v-model="queryType" @change="queryTypeChanged">
      <option value="0">基础类型</option>
      <option value="1">传奇</option>
      <option value="2">宝石</option>
    </select>
    <input type="text" id="inputText" list="matchedList" v-model="input" @input="inputing">
    <datalist id="matchedList">
      <option v-for="item of matchedList">{{ item }}</option>
    </datalist>
    <p>{{ output }}</p>
  </div>
</template>

<style>
.container>*:nth-child(n+2) {
  margin-left: 5px;
}
</style>
