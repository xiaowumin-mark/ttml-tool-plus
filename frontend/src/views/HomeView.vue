<script setup>
import Button from '@/components/ui/button.vue';
import ComboBox from '@/components/ui/combobox.vue';
import EditBox from '@/components/ui/edit-box.vue';
import icon from '@/components/ui/icon.vue';
import linkButton from '@/components/ui/link-button.vue';
import ListItem from '@/components/ui/list-item.vue';
import Radio from '@/components/ui/radio.vue';
import RadioGroup from '@/components/ui/radio-group.vue';
import Checkbox from '@/components/ui/checkbox.vue';
import CheckboxGroup from '@/components/ui/checkbox-group.vue';
import { ref, watch, onMounted, } from 'vue';
import { GetMe, GetRepo, GetUser } from '../../bindings/ttml-tool-plus/github-api/githubapiservice';
import { formatBytes, formatDateTime, formatTime } from '@/assets/tool';
const repoData = ref({});
const isLoading = ref(true);
GetRepo().then(res => {
  console.log(res);
  repoData.value = res;
  isLoading.value = false;
})

GetMe().then(res => {
  console.log(res);
})
</script>

<template>
  <main style="max-width: 1000px;" class="padding-view scroll mainview">
    <p style="font-size: 28px;">👋Hi！ 欢迎来到ttml-tool-plus</p>
    <div class="ds" v-if="!isLoading">
      <div class="basic item">
        <div class="title">仓库大小</div>
        <div class="text">
          <div>{{ formatBytes(repoData.size * 1024) }}</div>
        </div>
      </div>
      <div class="basic item">
        <div class="title">Star数量</div>
        <div class="text">
          <div>{{ repoData.stargazers_count }}</div>
        </div>
      </div>
      <div class="basic item">
        <div class="title">Fork数量</div>
        <div class="text">
          <div>{{ repoData.forks_count }}</div>
        </div>
      </div>
      <div class="basic item">
        <div class="title">更新时间</div>
        <div class="text">
          <div>{{ formatTime(repoData.updated_at) }}</div>
        </div>
      </div>
    </div>
  </main>
</template>



<style lang="css" scoped>
@import url("../components/ui/basic.css");


.mainview {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  gap: 20px;
}

.ds {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  gap: 5px;
  flex-wrap: wrap;
  width: 100%;
}

.ds>.item {
  /* 改为一行两个 */
  flex: 0 0 calc(50% - 5px);
  max-width: calc(50% - 5px);
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  gap: 8px;
  padding: 8px;
  box-sizing: border-box;
}

.ds>.item>.title {
  font-size: 16px;
  font-weight: bold;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}

.ds>.item>.title::before {
  content: "";
  width: 3px;
  height: 16px;
  display: inline-block;
  background-color: var(--user-color);
  margin-right: 5px;
  border-radius: 2px;
}

.ds>.item>.text {
  font-size: 20px;
  font-weight: bold;
}
</style>