import { message as Message } from "ant-design-vue";

const checkStatus = (status?: number) => {
  // 处理 HTTP 网络错误
  let message = "";
  switch (status) {
    case 401:
      message = "token 失效，请重新登录";
      // 这里可以触发退出的 action
      break;
    case 403:
      message = "拒绝访问";
      break;
    case 404:
      message = "请求地址错误";
      break;
    case 500:
      message = "服务器故障";
      break;
    default:
      message = "网络连接故障";
  }
  Message.error(message);
};

export default checkStatus;
