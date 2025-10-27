<template>
    <main @mousedown="md" @mouseup="mu" @mouseleave="ml">
        <icon :icon="props.icon" ref="sif"></icon>
    </main>
</template>
<script setup>
import { ref, defineProps, onMounted } from 'vue'
import icon from './icon.vue'
const sif = ref(null)
const props = defineProps({
    icon: {
        type: String,
    },
    montion: {
        type: String,
        default: 'none'
    }
})

const montions = {
    GlobalNavButton: {
        enter: [[
            {}, {
                transform: 'scaleX(0.5)'
            }
        ], {
            duration: 100,
            ease: 'easeInOut',
            fill: 'forwards'
        }],
        leave: [
            [{}, {
                transform: 'scaleX(1)'
            }], {
                duration: 100,
                ease: 'easeInOut',
                fill: 'forwards'
            }
        ]
    }
}

onMounted(() => {
    console.log(props.montion);
    console.log(montions[props.montion]);



})
const md = () => {
    if (props.montion == 'none') return;
    add(montions[props.montion].enter)
    console.log("enter");

}
const mu = () => {
    if (props.montion == 'none') return;
    //sif.value.animate(montions[props.montion].leave[0], montions[props.montion].leave[1])
    add(montions[props.montion].leave)
    console.log("leave");
}
const ml = () => {
    //sif.value.animate(montions[props.montion].leave[0], montions[props.montion].leave[1])
    //add(montions[props.montion].leave)
    //console.log("leave");
}

const add = (r) => {
    sif.value.icon.animate(r[0], r[1])

}
</script>

<style scoped>
main {
    width: 40px;
    height: 40px;
    max-height: 40px;
    max-width: 40px;
    border-radius: 4px;

    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;

    transition: background-color 0.1s ease-in-out,filter 0.1s ease-in-out,opacity 0.1s ease-in-out;

}

main:hover {

    background-color: rgba(104, 104, 104, 0.1);
}


main:active {
    background-color: rgba(104, 104, 104, 0.2);
    filter: brightness(0.8);
    opacity: 0.8;
}
</style>