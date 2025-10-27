// dialog-service.js
import { createVNode, render, nextTick, h } from 'vue'
import SmartDialog from '../ui/dialog.vue'
import Button from '../ui/button.vue'

export function dialog(options = {}) {
  const container = document.createElement('div')
  document.body.appendChild(container)

  // 真正销毁 DOM 的函数（在 after-leave 调用）
  const destroy = () => {
    render(null, container)
    if (container.parentNode) container.parentNode.removeChild(container)
  }

  // 创建 vnode，但初始为关闭状态（这样能触发 enter 动画）
  const vnode = createVNode(SmartDialog, {
    modelValue: false,
    title: options.headline || options.title || '',
    escClose: options.escClose ?? true,
    maskClose: options.maskClose ?? true,

    // 当组件内部触发 "close"（比如用户点按钮或 mask 点击）时
    onClose() {
      // 不要马上移除 DOM，改为触发离场动画
      if (vnode.component) {
        vnode.component.props.modelValue = false
      }
    },

    // transition 的结束事件：离场动画完成后销毁
    onAfterLeave() {
      destroy()
      options.onAfterClose?.()
    }
  }, {
    default: () => options.description || options.content || '',
    footer: () => {
      // 使用 JSX-like vnode children won't work here directly if you don't compile,
      // but createVNode accepts render functions — for simplicity create plain DOM buttons using vnode children array
      return (options.actions || []).map((act, idx) => {
        // 用 createVNode 创建按钮节点（原生 button）
        const onClick = () => {
          try {
            const ret = act.onClick?.()
            // 只有当回调显式返回 false 时阻止自动关闭
            if (ret !== false) {
              // 触发组件 close（组件会把 modelValue 设为 false，进而播放离场动画）
              if (vnode.component) vnode.component.props.modelValue = false
            }
          } catch (e) {
            // 若回调抛错，仍关闭并把错误抛给控制台
            console.error(e)
            if (vnode.component) vnode.component.props.modelValue = false
          }
        }
        // 返回一个简单 vnode：button
        /*return createVNode('button', {
          key: idx,
          onClick,
          style: { marginLeft: idx === 0 ? '0px' : '10px' }
        }, act.text || `action-${idx}`)*/
        return h(Button, {
          onClick: onClick
        }, act.text)
      })
    }
  })

  // 渲染 vnode（还是 modelValue: false）
  render(vnode, container)

  // 下一帧再打开，保证 enter 动画发生
  nextTick(() => {
    if (vnode.component) {
      vnode.component.props.modelValue = true
    }
  })

  // 返回一个可编程接口（比如外部可以直接关闭）
  return {
    close: () => {
      if (vnode.component) vnode.component.props.modelValue = false
    },
    destroy: () => destroy()
  }
}
