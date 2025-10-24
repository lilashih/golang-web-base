export default {
    success: (title, options = {}) => {
        ElementPlus.ElNotification.success({
            title,
            customClass: 'notification-success',
            ...options,
        })
    },
    error: (title, options = {}) => {
        ElementPlus.ElNotification.error({
            title,
            customClass: 'notification-error',
            ...options,
        })
    },
    warning: (title, options = {}) => {
        ElementPlus.ElNotification.warning({
            title,
            customClass: 'notification-warning',
            ...options,
        })
    },
    info: (title, options = {}) => {
        ElementPlus.ElNotification.info({
            title,
            customClass: 'notification-info',
            ...options,
        })
    },
};