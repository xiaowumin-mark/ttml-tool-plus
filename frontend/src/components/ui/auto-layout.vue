<template>
  <div class="layout" :style="{ height: h + 'px' }">
    <!-- 向 slot 暴露一个回调函数 onLayout -->
    <slot :onLayout="onLayout"></slot>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
const elements = ref([])
const h = ref(0)
const gap = 6
function onLayout(payload) {
  // 检查是否包含payload.item
  for (let i = 0; i < elements.value.length; i++) {
    if (elements.value[i].item === payload.item) {
      elements.value[i] = payload
      setTimeout(() => {
        setLayout(elements.value, {
          isinit: false,
          position: i
        })
      }, 150 * !elements.value[i].open)
      return
    }
  }
  elements.value.push(payload)

}
onMounted(() => {
  console.log('✅ auto-layout 组件已挂载')
  console.log(elements.value);
  setLayout(elements.value)

})
const setLayout = (list, options = {
  isinit: true,
  position: 0
}) => {
  let last = 10
  list.forEach((item, index) => {
    // 设置--position



    const t = (index - options.position + 1) * 20 * !options.isinit
    console.log(index, t);

    setTimeout(() => {
      item.item.style.setProperty('--position', last + "px")

      last += item.height + gap
      if (index === list.length - 1) {
        h.value = last
      }
    }, t)

  })
}
</script>

<style scoped>
.layout {
  width: 100%;
  overflow-y: hidden;
  transition: height 400ms cubic-bezier(0.190, 1.000, 0.220, 1.000);
}
</style>
