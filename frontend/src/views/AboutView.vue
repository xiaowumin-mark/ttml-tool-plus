<template>
    <main style="max-width: 1400px" class="padding-view scroll">
        <Dialog v-model="showDialog" title="Save your work?" :mask-close="false">
            <template #trigger>
                <Button>打开弹窗</Button>
            </template>
            <p>这是对话框中的文本。</p>
            <RadioGroup>
                <Radio label="1" value="1">1</Radio>
                <Radio label="2" value="2">2</Radio>
                <Radio label="3" value="3">3</Radio>
            </RadioGroup>
            <CheckboxGroup>
                <Checkbox label="1" value="1">1</Checkbox>
                <Checkbox label="2" value="2">2</Checkbox>
                <Checkbox label="3" value="3">3</Checkbox>
            </CheckboxGroup>
            <template #footer>
                <Button @click="showDialog = false">知道了</Button>
            </template>
        </Dialog>
        
        <Button @click="apidialog">
            打开API Dialog
        </Button>
    </main>
</template>
<script setup>
import { ref } from "vue";
import Dialog from "../components/ui/dialog.vue";
import Button from "../components/ui/button.vue";
import RadioGroup from "@/components/ui/radio-group.vue";
import Radio from "@/components/ui/radio.vue";
import CheckboxGroup from "@/components/ui/checkbox-group.vue";
import Checkbox from "@/components/ui/checkbox.vue";
import { dialog } from "@/components/api/dialog";
const showDialog = ref(false);
function apidialog() {
    dialog({
        headline: "安全警告 ⚠️",
        description: "确定要这么做吗？危险捏！",
        actions: [
            { text: "取消" },
            {
                text: "OK",
                onClick: () => {
                    console.log("勇士确认了任务")
                    return false  // 留在现场观察
                }
            }
        ]
    })
}
</script>

<style scoped>
@import url("../components/ui/basic.css");

main {
    width: 100%;
    height: 100%;
}
</style>
