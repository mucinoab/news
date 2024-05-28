import SendButton from './SendButton';

import { Card } from 'antd';

export interface Newsletter { id: number, name: string, description: string, subject: string, content: string }

const NewsletterItem = ({ newsletter }: { newsletter: Newsletter }) => {
  return <Card title={newsletter.name} extra={<SendButton id={newsletter.id} />} >
    <p>{newsletter.description}</p>
    <Card title={`Subject: ${newsletter.subject}`} type="inner" >
      <a href={`/file/${newsletter.content}`} target="_blank"> See attachment to send </a>
    </Card>
  </Card>
};

export default NewsletterItem;
