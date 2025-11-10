import config from "../config/config.js";
import loading from "./loading.js";
import notify from "./notify.js";

export default {
    // ajax請求
    fetch: function (method, url, data = null) {
        method = method.toLowerCase();

        let headers = {
            "Content-Type": "application/json",
            Accept: "application/json",
        };
        if (method === "get") {
            let params = new URLSearchParams(data);
            url += `?${params.toString()}`;
            data = null;
        } else if (data instanceof FormData) {
            // 上傳檔案
            data.append("_method", method);
            method = "post";
            delete headers["Content-Type"];
        } else if (typeof data === typeof {}) {
            data = JSON.stringify(data);
        }

        // credentials for enable cookie
        return fetch(url, {
            method,
            body: data,
            headers,
            credentials: "include",
        }).then(async (response) => {
            if (response.status === 200) {
                return await response.json();
            } else {
                console.log("fetch error response", response);
                let err = {
                    code: response.status,
                    message: `${response.statusText}: ${response.url}`,
                    data: {},
                };
                let res = await response.text();
                try {
                    const resData = JSON.parse(res);
                    err.data = resData.data || {};
                    err.message = resData.message || err.message;
                } catch {
                    err.data = res;
                }

                throw err;
            }
        });
    },

    // 常用的api請求
    api: function (method, url, data = {}) {
        loading.show();

        return this.fetch(method, `${config.url.api}/${url}`, data)
            .then((result) => {
                return result;
            })
            .catch((err) => {
                if (err.code != 422) {
                    console.log("api error", err);
                    err.message = this.isEmpty(err.message) ? err.exception : err.message;
                    notify.error(err.message);
                } else {
                    if (Array.isArray(err.data.errors)) {
                        err.data.errors = this.flattenErrors(err.data.errors);
                    }
                }
                throw err;
            })
            .finally(function () {
                loading.hide();
            });
    },

    flattenErrors: (errors) => {
        errors = errors.reduce((acc, curr) => {
            acc[curr.name] = curr.message;
            return acc;
        }, {});
        return errors;
    },

    // 檢查是否為空
    isEmpty: (target) => {
        if (typeof target === "object") {
            for (let key in target) {
                if (target.hasOwnProperty(key)) return false;
            }
            return true;
        } else if (target === undefined || target === null || target === "" || target === "undefined") {
            return true;
        }
        return false;
    },

    // 深拷貝
    deepCopy: (obj) => {
        return JSON.parse(JSON.stringify(obj));
    },

    scrollToTop: (selector = `main`) => {
        document.querySelector(selector).scrollTo({
            top: 0,
            behavior: "smooth",
        });
    },
};