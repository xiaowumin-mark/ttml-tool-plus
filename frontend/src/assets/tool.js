function formatTime(input) {
    // 判断是否为纯数字（毫秒）
    if (typeof input === 'number' || /^\d+$/.test(input)) {
        const ms = Number(input);
        if (ms < 1000 * 60 * 60 * 24) {
            // 小于一天，视为时长
            return formatDuration(ms);
        }
        // 大于一天，可能是时间戳
        const date = new Date(ms);
        return formatDateTime(date);
    }

    // 解析时间字符串
    const date = new Date(input);
    if (isNaN(date)) return '无效时间';

    return formatDateTime(date);
}

function formatDuration(ms) {
    const totalSeconds = Math.floor(ms / 1000);
    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;

    if (minutes === 0) return `${seconds}秒`;
    return `${minutes}分${seconds.toString().padStart(2, '0')}秒`;
}

function formatDateTime(date) {
    const now = new Date();
    const diff = (now - date) / 1000; // 秒差
    const absDiff = Math.abs(diff);

    const pad = n => n.toString().padStart(2, '0');
    const hours = pad(date.getHours());
    const minutes = pad(date.getMinutes());

    // 相对时间简化显示
    if (absDiff < 60) {
        return '刚刚';
    } else if (absDiff < 3600) {
        const mins = Math.floor(absDiff / 60);
        return `${mins} 分钟前`;
    } else if (absDiff < 86400 && date.getDate() === now.getDate()) {
        const hoursAgo = Math.floor(absDiff / 3600);
        return `${hoursAgo} 小时前`;
    }

    // 昨天
    const yesterday = new Date(now);
    yesterday.setDate(now.getDate() - 1);
    if (
        date.getDate() === yesterday.getDate() &&
        date.getMonth() === yesterday.getMonth() &&
        date.getFullYear() === yesterday.getFullYear()
    ) {
        return `昨天 ${hours}:${minutes}`;
    }

    // 同一年
    if (date.getFullYear() === now.getFullYear()) {
        return `${date.getMonth() + 1}月${date.getDate()}日 ${hours}:${minutes}`;
    }

    // 不同年份
    return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日 ${hours}:${minutes}`;
}

function formatBytes(bytes, decimals = 2) {
    if (!bytes) return '0 B';
    
    const k = 1024;
    const units = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB'];

    const i = Math.floor(Math.log(bytes) / Math.log(k));
    const value = (bytes / Math.pow(k, i)).toFixed(decimals);

    return `${value} ${units[i]}`;
}


export { formatTime, formatDuration, formatDateTime ,formatBytes};
