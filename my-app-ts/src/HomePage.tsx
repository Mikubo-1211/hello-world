import React, { useState, useEffect } from 'react';
import { useCookies } from 'react-cookie';

type HomePageProps = {
  handleLogout: () => void;
};

type Message = {
  id: string;
  edit: string;
  user_id: string;
  user_name: string;
  content: string;
  created: string;
};

type Channel = {
  channel_id: string;
  channel_name: string;
  descript: string;
};

type User = {
  user_id: string;
  user_name: string;
  user_password: string;
  user_email: string;
};

const HomePage: React.FC<HomePageProps> = ({ handleLogout }) => {
  const [message, setMessage] = useState('');
  const [messages, setMessages] = useState<Message[]>([]);
  const [channels, setChannels] = useState<Channel[]>([]);
  const [user, setUser] = useState<User | null>(null);
  const [cookies] = useCookies(['userEmail']);
  const userEmail = cookies.userEmail || '';
  const [currentChannel, setCurrentChannel] = useState<Channel | null>(null);
  const [editMessageId, setEditMessageId] = useState('');
  const [editMessage, setEditMessage] = useState('');

  useEffect(() => {
    fetchChannel();
  }, []);

  useEffect(() => {
    fetchUsers(userEmail);
  }, [userEmail]);

  const fetchMessage = async (channelId: string) => {
    try {
      const result = await fetch(`https://hello-world-2-xyex4gyyzq-uc.a.run.app/channel?channel_id=${channelId}`, {
        method: 'GET',
      });
      if (!result.ok) {
        throw Error(`Failed to fetch messages: ${result.status}`);
      }
      const messagesData = await result.json();
      setMessages(messagesData);
    } catch (error) {
      console.error(error);
    }
  };

  const fetchUsers = async (email: string) => {
    try {
      const result = await fetch(`https://hello-world-2-xyex4gyyzq-uc.a.run.app/users?user_email=${email}`, {
        method: 'GET',
      });
      if (!result.ok) {
        throw Error(`Failed to fetch user: ${result.status}`);
      }
      const userData = await result.json();
      setUser(userData);
    } catch (error) {
      console.error(error);
    }
  };

  const fetchChannel = async () => {
    try {
      const result = await fetch('https://hello-world-2-xyex4gyyzq-uc.a.run.app/channels', {
        method: 'GET',
      });
      if (!result.ok) {
        throw Error(`Failed to fetch channels: ${result.status}`);
      }
      const channelsData = await result.json();
      setChannels(channelsData);
    } catch (error) {
      console.error(error);
    }
  };

  const handleSendMessage = async () => {
    if (!message) {
      alert('メッセージを入力してください');
      return;
    }
    if (!currentChannel) {
      alert('チャンネルを選択してくださいな');
      return;
    }
    if (!user) {
      alert('ユーザーIDがありません');
      return;
    }
    if (message.length > 150) {
      alert('メッセージは150文字以内で入力してください');
      return;
    }
    try {
      const url = new URL('https://hello-world-2-xyex4gyyzq-uc.a.run.app/message');
      url.searchParams.append('user_id', user.user_id);
      url.searchParams.append('channel_id', currentChannel.channel_id);
      url.searchParams.append('content', message);

      const result = await fetch(url.toString(), {
        method: 'POST',
      });
      if (!result.ok) {
        throw Error(`Failed to send message: ${result.status}`);
      }

      setMessage('');
      fetchMessage(currentChannel.channel_id);
    } catch (error) {
      console.error(error);
    }
  };

  const handleEditMessage = (messageId: string, messageContent: string) => {
    setEditMessageId(messageId);
    setEditMessage(messageContent);
  };

  const handleUpdateMessage = async () => {
    try {
      const response = await fetch(`https://hello-world-2-xyex4gyyzq-uc.a.run.app/message`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id: editMessageId, content: editMessage }),
      });

      if (!response.ok) {
        throw new Error(`Failed to update message: ${response.status}`);
      }

      setEditMessageId('');
      setEditMessage('');
      fetchMessage(currentChannel?.channel_id || '');
    } catch (error) {
      console.error(error);
    }
  };

  const handleDeleteMessage = async (id: string) => {
    const confirmed = window.confirm('本当にメッセージを削除してもよろしいですか？');
    if (!currentChannel) {
      alert('チャンネルを選択してくださいな');
      return;
    }
    if (confirmed) {
      try {
        const response = await fetch(`https://hello-world-2-xyex4gyyzq-uc.a.run.app/message?id=${id}`, {
          method: 'DELETE',
        });
        if (!response.ok) {
          throw new Error(`Failed to delete message: ${response.status}`);
        }
        fetchMessage(currentChannel.channel_id);
      } catch (error) {
        console.error(error);
      }
    }
  };

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(event.target.value);
  };

  const handleEditInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEditMessage(event.target.value);
  };

  return (
    <div>
      {user && (
        <>
          <label>ユーザー名: {user.user_name}</label>
        </>
      )}

      <h3>チャンネルリスト</h3>
      <ul>
        {channels.map((channel) => (
          <li key={channel.channel_id}>
            <button
              onClick={() => {
                fetchMessage(channel.channel_id);
                setCurrentChannel(channel);
              }}
            >
              {channel.channel_name}
            </button>
          </li>
        ))}
      </ul>

      <h3>メッセージ</h3>
      <ul>
      {messages.map((message) => (
  <li key={message.id}>
    <div>
      <span>名前: {message.user_name}</span>
      <br />
      {editMessageId === message.id ? (
        <div>
          <input type="text" value={editMessage} onChange={handleEditInputChange} />
          <button onClick={handleUpdateMessage}>更新</button>
        </div>
      ) : (
        <>
          <span>メッセージ: {message.content}</span>
          {message.edit === '1' && <span> (編集済)</span>}
          
          <br />
          {message.user_id === user?.user_id && (
            <div>
              <button onClick={() => handleDeleteMessage(message.id)}>削除</button>
              <button onClick={() => handleEditMessage(message.id, message.content)}>編集</button>
            </div>
          )}
        </>
      )}
    </div>
  </li>
))}

      </ul>

      <div>
        <input type="text" value={message} onChange={handleInputChange} />
        <button onClick={handleSendMessage}>送信</button>
      </div>
      <button onClick={handleLogout}>ログアウト</button>
    </div>
  );
};

export default HomePage;

