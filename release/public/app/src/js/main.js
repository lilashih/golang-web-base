import menu from "./menu/menu.js";
import Dashboard from "../view/Dashboard.js";

const router = createRouter({
    history: createWebHistory(),
    routes: menu.flatten(menu.dashboard),
});

const template = /*html*/ `
<Dashboard :menus="menus"/>
`;

// 建立 Vue 實例
const app = createApp({
    components: {
      Dashboard
    },
    template,
    setup() {
        return {
            ...ElementPlus,
            menus: menu.dashboard,
        };
    },
    watch: {
        $route: {
            immediate: true,
            handler(to, from) {
                document.title = to.meta.header || 'Dashboard';
            }
        },
    },
});

app.config.errorHandler = (err, vm, info) => {
    console.error("Error:", err);
    console.info("Error VueComponent:", vm);
    console.log("Error info:", info);
};

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}

app.use(router).use(ElementPlus, { locale: ElementPlusLocaleZhTw }).mount("#app");
