<template>
    <main style="width: 100%;" class="padding-view scroll">
        <div style="max-width: 1000px;;position: relative;margin: 0 auto;">
            <p style="font-size: 28px;">设置</p>
            <div v-if="!isLogin">
                <p style="line-height: 40px;">您貌似并未登录github，github api将限制每小时调用60次API，点击下方按钮开始登录🙌</p>
                <Button style="width: 150px;" @click="OpenOauthWindow">登录</Button>
            </div>
            <AutoLayout v-slot="{ onLayout }" style="margin-top: 20px;">
                <!--<AutoLayoutItem @layout="onLayout" icon="&#xE77B;" title="Github 用户名" desc="可为该用户显示更多信息">
                    <template #end>
                        <EditBox placeholder="输入github用户名" @change="configChange('githubUser', $event)"
                            :value="config['githubUser']"></EditBox>
                    </template>
</AutoLayoutItem>-->

                <AutoLayoutItem @layout="onLayout" icon="&#xF3B1;" title="退出登录" desc="退出登录后，将无法同步数据" v-if="isLogin">
                    <template #end>
                        <Button @click="LogOutUser">退出登录</Button>
                    </template>
                </AutoLayoutItem>

            </AutoLayout>


        </div>
    </main>
</template>
<script setup>
import AutoLayoutItem from '@/components/ui/auto-layout-item.vue';
import AutoLayout from '@/components/ui/auto-layout.vue';
import Button from '@/components/ui/button.vue';
import { onMounted, onUnmounted, ref, watch } from 'vue';
import EditBox from '@/components/ui/edit-box.vue';
import { ChangeConfigAndSave, GetConfig, IsLogin, Logout } from '../../bindings/ttml-tool-plus/config/configapiservice';
import { OpenOauthWindow } from '../../bindings/ttml-tool-plus/github-api/githubapiservice';
import { Events, Window } from '@wailsio/runtime';
const isLogin = ref(false);
const config = ref({});
GetConfig().then(res => {
    config.value = res;
})
onMounted(() => {
})
IsLogin().then(res => {
    isLogin.value = res;
})
const configChange = (key, val) => {
    console.log(key, val);
    ChangeConfigAndSave(key, val).then(res => {
        config.value = res;
    })
}

Events.On("oauth_error", (data) => {
    console.log("oauth出现错误", data);

})
Events.On("oauth_success", (data) => {
    console.log("oauth成功", data);
    IsLogin().then(res => {
        isLogin.value = res;
        console.log("oauth成功",res);
        
    })
})
Events.On("oauth_stopped", (data) => {
    console.log("oauth取消", data);
})
Events.On("oauth_user_cancel", (data) => {
    console.log("oauth用户取消", data.data[0]);
})
onUnmounted(() => {
    Events.Off("oauth_error");
    Events.Off("oauth_success");
    Events.Off("oauth_stopped");
    Events.Off("oauth_user_cancel");
})

const LogOutUser = () => {
    Logout().then(res => {
        isLogin.value = false;
    })

}

</script>

<style lang="css" scoped>
@import url("../components/ui/basic.css");
</style>