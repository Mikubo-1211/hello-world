import React, { useState, ChangeEvent, useContext } from 'react';
import { createUserWithEmailAndPassword } from "firebase/auth";
import { fireAuth } from "./firebase";
import { Link } from 'react-router-dom';
import { UserContext } from './UserContext';

interface SignUpFormProps {
  handleLogin: () => void; // handleLogin プロパティの型定義
}

const SignUpForm: React.FC<SignUpFormProps> = ({ handleLogin }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [username, setUsername] = useState('');
  const { setUserEmail } = useContext(UserContext); // UserContextからsetUserEmailを取得

  const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const handleUsernameChange = (event: ChangeEvent<HTMLInputElement>) => {
    setUsername(event.target.value);
  };

  const signUpWithEmail = async (): Promise<void> => {
    if (!username || username.length === 0) {
      alert("名前を入力してください。");
      return;
    }

    if (username.length > 50) {
      alert("名前は50文字以下で入力してください。");
      return;
    }

    if (email.length > 50) {
      alert("メールアドレスは50文字以下で入力してください。");
      return;
    }
    

    try {
      const res = await createUserWithEmailAndPassword(fireAuth, email, password);
      const user = res.user;
      alert("新規登録ユーザー: " + username); // ユーザー名を表示する
      handleLogin(); // ログイン成功時に handleLogin を呼び出す
      setUserEmail(email); // userEmailを更新

      if (!username) {
        alert("Please enter name");
        return;
      }

      try {
        const result = await fetch("https://hello-world-2-xyex4gyyzq-uc.a.run.app/users", {
          method: "POST",
          body: JSON.stringify({
            user_name: username,
            user_email: email,
            user_password: password
          }),
        });
        if (!result.ok) {
          throw Error(`Failed to create user: ${result.status}`);
        }
      } catch (err) {
        const errorMessage = (err as Error).message; // エラーオブジェクトの型を明示的に指定する
        alert(errorMessage);
      }
    } catch (err) {
      const errorMessage = (err as Error).message; // エラーオブジェクトの型を明示的に指定する
      alert(errorMessage);
    }
  };

  return (
    <div>
      <h2>新規登録</h2>
      <label>
        名前:
        <input type="username" value={username} onChange={handleUsernameChange} />
      </label>
      <br />
      <label>
        メールアドレス:
        <input type="email" value={email} onChange={handleEmailChange} />
      </label>
      <br />
      <label>
        パスワード:
        <input type="password" value={password} onChange={handlePasswordChange} />
      </label>
      <br />
      <button onClick={signUpWithEmail}>
        新規登録
      </button>
    </div>
  );
};

export default SignUpForm;
