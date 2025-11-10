import config from "../js/config/config.js";

const template = /*html*/ `
<el-container class="h-lvh">

    <!-- 左側菜單 -->
    <el-aside class="menu-aside" :class="{ 'collapsed': isCollapse }">

      <el-menu class="el-menu-vertical-demo" :collapse="isCollapse" :router="false">
        <!-- 收合菜單 -->
        <el-menu-item class="menu-collapse">
          <el-icon v-show="isCollapse" @click="setCollapse(false)"><ArrowRight /></el-icon>
          <el-icon v-show="!isCollapse" @click="setCollapse(true)"><ArrowLeft /></el-icon>
          <span>{{ name }}</span>
        </el-menu-item>

        <template v-for="(menu, index) in menus" :key="index">
            <!-- 多層菜單 -->
            <el-sub-menu v-if="menu.hasOwnProperty('children')" :index="getIndex(index)">
                <template #title>
                  <el-icon><component :is="menu.meta.icon" /></el-icon>
                  <span>{{ menu.meta.title }}</span>
                </template>
                <el-menu-item-group>
                    <router-link v-for="(child, cIndex) in menu.children" :to="{ name: child.name }" :key="getIndex(index, cIndex)">
                        <el-menu-item :index="getIndex(index, cIndex)">{{ child.meta.title }}</el-menu-item>
                    </router-link>
                </el-menu-item-group>
            </el-sub-menu>
            <!-- 單層菜單 -->
            <router-link v-else :to="{ name: menu.name }">
                <el-menu-item :index="getIndex(index)">
                  <el-icon><component :is="menu.meta.icon" /></el-icon>
                  <span>{{ menu.meta.title }}</span>
                </el-menu-item>
            </router-link>
        </template>
      </el-menu>
    </el-aside>

    <!-- 右側內容 -->
    <router-view :key="$route.fullPath"></router-view>
</el-container>
`;

export default {
    props: {
        menus: {
            type: Array,
        },
    },
    setup(props) {
        const isCollapse = ref(false);
        const getIndex = (index1, index2 = -1) => `${index1}_${index2}`;
        const setCollapse = (isCollapseNew) => isCollapse.value = isCollapseNew
        const name = ref(config.name);

        return {
            name,
            isCollapse,
            getIndex,
            setCollapse,
        };
    },
    template,
};