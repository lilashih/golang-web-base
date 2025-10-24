import Home from "../../view/component/Home.js";
import config from "../config/config.js";
import helper from "../helper/helper.js";

const path = config.url.view;

let menus = localStorage.getItem("menus");

if (helper.isEmpty(menus)) {
    const data = await helper.api("get", "menus");
    const items = data.data.menus;

    menus = [
        {
            meta: {
                title: "首頁",
                icon: "HomeFilled",
            },
            path: `${path}`,
            name: "index",
            component: Home,
        },
    ];

    items.map((item) => {
        let m = {
            meta: {
                header: item.group,
                title: item.group,
                icon: "StarFilled",
            },
            children: [],
        };

        item.children.map((item2) => {
            const isSettingMenu = item2.id.substring(0, 7) == 'setting';
            const c = isSettingMenu ? `Setting` : item2.id[0].toUpperCase() + item2.id.substring(1);

            m.children.push({
                meta: {
                    header: `${item.group} / ${item2.name}`,
                    title: item2.name,
                    description: item2.description,
                    icon: "",
                    api: item2.path,
                    id: item2.id,
                },
                path: `${path}/${item2.id}`,
                name: item2.id,
                component: () => import(`../../view/component/${c}.js`),
            });
        });

        menus.push(m);
    });
    // localStorage.setItem("menus", JSON.stringify(menus));
} else {
    menus = JSON.parse(menus);
}

const flatten = (objects) => {
    let flattened = [];

    objects.forEach((obj) => {
        if (obj.children) {
            obj.children.forEach((group) => {
                flattened.push({
                    ...group,
                    title: group.title,
                });
            });
        } else {
            flattened.push(obj);
        }
    });

    return flattened;
};

export default {
    flatten,
    dashboard: menus,
};
