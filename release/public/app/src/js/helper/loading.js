export default {
    show: (options = {}) => {
        ElementPlus.ElLoading.service({
            text: '載入中，請稍後',
            background: 'rgba(0, 0, 0, 0.7)',
            lock: true,
            ...options,
        })
    },
    hide: () => {
        ElementPlus.ElLoading.service().close()
    },
};