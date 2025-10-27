<template>
    <div class="auto-layout-item" ref="mainitem" :class="{ open: open }">
        <div class="main-item  basic" :class="{ canactive: props.toggle, disabled: props.disabled }" ref="itemele"
            @click="props.toggle ? open = !open : null">
            <div class="icon" v-show="props.icon">
                <Icon :icon="props.icon" size="20"></Icon>
            </div>
            <div class="cont">
                <div class="title">
                    {{ props.title }}
                </div>
                <div class="desc">
                    {{ props.desc }}
                </div>
            </div>
            <div class="active-icon">
                <slot name="end"></slot>
            </div>
            <div class="toggle-icon" v-show="props.toggle">
                <Icon icon="&#xE70D;" size="14"></Icon>
            </div>
        </div>
        <div class="toggle"  v-if="props.toggle" ref="toggleele">
            <div class="cont">
                <slot name="toggle">

                </slot>
            </div>
        </div>
    </div>
</template>

<script setup>

import { ref, onMounted, defineProps, watch } from "vue"
import Icon from "./icon.vue"
const props = defineProps({
    toggle: {
        type: Boolean,
        default: false
    },
    disabled: {
        type: Boolean,
        default: false
    },
    icon: {
        type: String,
        default: ""
    },
    title: {
        type: String,
        default: ""
    },
    desc: {
        type: String,
        default: ""
    }
})
const open = defineModel(false)
watch(open, (newVal) => {
    if (newVal) {
        emit('layout', { item: mainitem.value, height: mainitem.value.offsetHeight ,open:open.value})
    } else {
        emit('layout', { item: mainitem.value, height: itemele.value.offsetHeight ,open:open.value})
    }
})
const mainitem = ref(null)
const toggleele = ref(null)
const itemele = ref(null)
const emit = defineEmits(['layout'])

onMounted(() => {
    emit('layout', { item: mainitem.value, height: itemele.value.offsetHeight ,open:open.value})

})
</script>

<style scoped>
@import url("./basic.css");

.auto-layout-item {

    position: absolute;
    transform: translateY(var(--position));
    width: 100%;
    box-sizing: border-box;

    overflow: hidden;
    transition: transform 400ms cubic-bezier(0.190, 1.000, 0.220, 1.000);
}

.main-item {
    position: relative;
    width: 100%;
    height: 66px;
    box-sizing: border-box;
    z-index: 2;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    padding-right: 10px;
}
.open>.main-item{
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
}

.main-item>.icon {
    width: 60px;
    height: 48px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.main-item>.cont {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
    box-sizing: border-box;

}

.main-item>.cont>.title {
    font-size: 14px;
}

.main-item>.cont>.desc {
    font-size: 12px;
    opacity: 0.8;
}

.main-item>.active-icon {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.main-item>.toggle-icon {
    width: 24px;
    height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: transform 0.4s cubic-bezier(0.190, 1.000, 0.220, 1.000);
    transition-delay: 150ms;
}

.auto-layout-item>.toggle {
    position: relative;
    z-index: 1;
    width: 100%;
    overflow: hidden;
    border-bottom-left-radius:5px;
    border-bottom-right-radius: 5px;

}

.open>.toggle>.cont {
    transform: translateY(0);
}
.open>.main-item>.toggle-icon {
    transform: rotate(180deg);
}

.toggle>.cont {

    transition: transform 0.4s cubic-bezier(0.190, 1.000, 0.220, 1.000);
    transform: translateY(calc(-100% - 3px));
    transform-origin: top;
}
</style>
