<template>
    <main class="scroll" style="--wails-draggable: drag;">
        <div class="mainview">
            <p>正在进行授权中，关闭则取消授权</p>
            <p>倒计时：{{ countdown }}秒</p>
            <Button @click="StopOAuth" style="color:red" class="noDrag" full :disabled="isOnOAuth">取消授权</Button>
        </div>
        <div class="stop" :class="{ show: isOnOAuth }">
            正在关闭服务<br/>请勿关闭窗口<br/>这可能需要10秒左右的时间
        </div>
    </main>
</template>
<script setup>
import { Application, Events, Window } from '@wailsio/runtime';
import { ref } from 'vue';
import Button from '@/components/ui/button.vue';
import { StartOAuth, StopOAuth } from '../../bindings/ttml-tool-plus/github-api/githubapiservice';
import { GetConfig } from '../../bindings/ttml-tool-plus/config/configapiservice';

const countdown = ref(0);
const isOnOAuth = ref(false);
Events.On("oauth_error", (data) => {
    console.log("oauth出现错误", data.data[0]);

})
Events.On("oauth_success", (data) => {
    console.log("oauth成功", data.data[0]);

})
Events.On("oauth_stopped", (data) => {
    console.log("oauth取消", data.data[0]);

})
Events.On("oauth_countdown", (data) => {
    console.log("oauth倒计时", data.data[0]);
    countdown.value = data.data[0];

})
Events.On("oauth_user_cancel", (data) => {
    console.log("oauth用户取消", data.data[0]);
})

Events.On("oauth_will_stopped", (data) => {
    isOnOAuth.value = true;
    console.log("开始关闭服务");

})
GetConfig().then(res => {
    StartOAuth(res['clientId']).then(res => {
        console.log("oauth开始", res);
    })
})
</script>
<style scoped>
@import url("../components/ui/basic.css");

main {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
}
.mainview { 
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 8px;

}

.noDrag {
    -webkit-app-region: no-drag;
}

.stop {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.8);
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
    font-size: 20px;
    font-weight: bold;
    backdrop-filter: blur(5px);
    opacity: 0;
    pointer-events: none;
    transition: opacity 0.2s ease-in-out;
    text-align: center;
}

.stop.show {
    opacity: 1;
    pointer-events: all;
}
</style>