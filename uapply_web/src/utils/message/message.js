import {
    createDiscreteApi,
    lightTheme
} from "naive-ui";
// createDiscreteApi传参
// 直接使用亮色主题
const configProviderPropsRef = {
    theme: lightTheme,
};

const { message, notification, dialog, loadingBar } = createDiscreteApi(["message", "dialog", "notification", "loadingBar"], {
    configProviderProps: configProviderPropsRef
});

export {
    message,
    notification,
    dialog,
    loadingBar
}