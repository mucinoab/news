import React from 'react';
import { useNavigate } from "react-router-dom";
import { Form, Flex, Input, Upload, Button, Typography, Divider } from 'antd';
import { UploadOutlined, QuestionCircleOutlined } from '@ant-design/icons';

const NewForm = () => {
  const navigate = useNavigate();
  const [form] = Form.useForm();

  const onFinish = (formData: any) => {
    formData["content"] = formData["content"]["file"]["name"];

    if (formData["recipients"]) { // The receipients file is optional
      formData["recipients"] = formData["recipients"]["file"]["name"];
    }

    fetch('/newsletter/create', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(formData),
    }).then(async (response) => {
      if (response.ok) {
        navigate("/");
      } else {
        throw new Error('Failed to submit form: ' + await response.text());
      }
    }).catch(console.error);
  };

  return <React.Fragment>
    <Flex justify="center">
      <Typography.Title level={1}>Create a Newsletter</Typography.Title>
    </Flex>

    <Divider />

    <Flex justify="center" style={{ height: "100vh" }}>
      <Form
        name="newsletter-form"
        labelCol={{ span: 8 }}
        wrapperCol={{ span: 16 }}
        onFinish={onFinish}
        form={form}
      >

        <Form.Item
          label="Name"
          name="name"
          rules={[{ required: true, message: "Please input a name for your newsletter" }]}>
          <Input />
        </Form.Item>

        <Form.Item label="Description" name="description" >
          <Input.TextArea rows={4} />
        </Form.Item>

        <Divider>
          <Typography.Title level={2}>Content</Typography.Title>
        </Divider>

        <Form.Item
          label="Subject"
          name="subject"
          rules={[{ required: true, message: "Please input a subject" }]}
          tooltip={{ title: "Write a nice subject line for your email", icon: <QuestionCircleOutlined /> }}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Content"
          name="content"
          rules={[{ required: true, message: "Please upload the main content" }]}
          tooltip={{ title: "Upload the main content of your newsletter (a PDF or PNG file). 32MB max", icon: <QuestionCircleOutlined /> }}
        >
          <Upload action="file/upload" accept=".pdf,.png" maxCount={1} >
            <Button icon={<UploadOutlined />}>Upload Content File</Button>
          </Upload>
        </Form.Item>

        <Divider>
          <Typography.Title level={2}>Recipients</Typography.Title>
        </Divider>

        <Form.Item
          label="Batch"
          name="recipients"
          tooltip={{ title: "Batch update your receipients, this can be a CSV file where each row contains an email or a TXT file where the emails are separated by commas. 32MB max", icon: <QuestionCircleOutlined /> }}
        >
          <Upload action="file/upload" accept=".csv,.txt" maxCount={1} >
            <Button icon={<UploadOutlined />}>Click to Upload File</Button>
          </Upload>
        </Form.Item>

        <Form.Item
          label="Recipient"
          name="recipient"
          rules={[{ type: "email", message: "Invalid email address" }]}
        >
          <Input placeholder="example@org.com" />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Button type="primary" htmlType="submit"> Create newsletter </Button>
        </Form.Item>
      </Form>
    </Flex>
  </React.Fragment >;
};

export default NewForm;
