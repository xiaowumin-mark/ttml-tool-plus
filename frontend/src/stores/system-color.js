import { defineStore } from "pinia"
import { GetWindowsThemeColor } from "../../bindings/ttml-tool-plus/system/systemapiservice"
import { ref } from "vue"

export const useColorStore = defineStore("color", () => {
    // 默认色：0, 120, 212
    const color = ref({
        r: 0,
        g: 120,
        b: 212
    })

    const setColor = (r, g, b) => {
        color.value.r = r
        color.value.g = g
        color.value.b = b
    }

    const getColor = () => color.value

    const getColorString = () =>
        `rgb(${color.value.r}, ${color.value.g}, ${color.value.b})`

    const applyColorToStyle = () => {
        document.documentElement.style.setProperty(
            "--user-color",
            getColorString()
        )
    }

    const setColorStyle = (r, g, b) => {
        setColor(
            r * 0.55 + 255 * 0.45,
            g * 0.55 + 255 * 0.45,
            b * 0.55 + 255 * 0.45
        )
        applyColorToStyle()
    }

    const loadSystemColor = async () => {
        try {
            const c = await GetWindowsThemeColor()
            console.log("SystemColor", c)
            setColorStyle(c.R, c.G, c.B)
        } catch (e) {
            console.log("Failed to load system color", e)
            applyColorToStyle() // 使用默认值
        }
    }

    // 初始化同步执行，仅使用默认颜色设置
    applyColorToStyle()
    // 异步获取系统主题色
    loadSystemColor()

    return {
        color,
        setColor,
        getColor,
        getColorString,
        setColorStyle,
        loadSystemColor
    }
})
