import { Button, Popconfirm, notification } from "antd";

type NotificationType = 'success' | 'error';

const SendButton = ({ id }: { id: number }) => {
  const [notificationApi, notificationContainer] = notification.useNotification();

  const openNotification = (message: string, type: NotificationType) => {
    notificationApi[type]({ message: "Distribution Result", description: message });
  };

  const sendNewsletter = () => {
    fetch(`/newsletter/${id}/send`, { method: 'POST' }).then(async response => {
      const message = response.ok ? 'Newsletter sent successfully' : await response.text();
      const type = response.ok ? 'success' : 'error';
      openNotification(message, type);
    })
  };

  return <Popconfirm
    title="Are you sure you want to SEND this newsletter?"
    onConfirm={sendNewsletter}
    okText="Yes"
    cancelText="No"
  >
    {notificationContainer}
    <Button style={{ marginLeft: 25 }}> Send </Button>
  </Popconfirm>;
}

export default SendButton;
