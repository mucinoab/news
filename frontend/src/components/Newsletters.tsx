import React, { useEffect, useState } from 'react';
import { Button, Card, Divider, Flex, FloatButton, List, Typography } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

import NewsletterItem, { Newsletter } from './NewsletterItem';

const Newsletters = () => {
  const navigate = useNavigate();
  const [news, setNews] = useState<Array<Newsletter> | undefined>(undefined);

  useEffect(() => {
    fetch('/newsletter/')
      .then(response => response.json())
      .then(data => setNews(data))
      .catch(console.error);
  }, []);

  let items = null;

  if (news) {
    items = news.map((newsletter: Newsletter) => <NewsletterItem key={newsletter.id} newsletter={newsletter} />);
  } else {
    const loadingCard = <Card loading extra={<Button >Send</Button >} style={{ width: "60vw" }} />;
    items = new Array(3).fill(loadingCard);
  }

  return <React.Fragment>
    <Flex justify="center">
      <Typography.Title level={1}>Your Newsletters</Typography.Title>
    </Flex>

    <Divider />
    <Flex justify="center" style={{ width: "100%" }}>
      <List
        grid={{ gutter: 16, column: 1 }}
        dataSource={items}
        renderItem={item => (<List.Item> {item} </List.Item>)}
      />
    </Flex>

    <FloatButton
      icon={<PlusOutlined />}
      type="primary"
      onClick={() => navigate("/new")}
    />
  </React.Fragment>;
};

export default Newsletters;
