export default {
    alert: (title, message = "", options = {}) => {
        return ElementPlus.ElMessageBox.alert(message, title, {
            confirmButtonText: "確認",
            ...options,
        });
    },
    confirm: (title, message = "", options = {}) => {
        return ElementPlus.ElMessageBox.confirm(message, title, {
            confirmButtonText: "確認",
            cancelButtonText: "取消",
            ...options,
        });
    },
    deleteConfirm: (message, options = {}) => {
        return ElementPlus.ElMessageBox.confirm(message, "刪除警告", {
            confirmButtonText: "刪除",
            cancelButtonText: "取消",
            type: 'error',
            icon: markRaw(ElementPlusIconsVue.Delete),
            ...options,
        });
    },
};
