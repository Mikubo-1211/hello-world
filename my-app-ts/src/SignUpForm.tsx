import React, { useState, ChangeEvent } from 'react';
import { createUserWithEmailAndPassword } from "firebase/auth";
import { fireAuth } from "./firebase";
import { Link } from 'react-router-dom';
import { useCookies } from 'react-cookie';
import { useEmailContext } from './Context';

interface SignUpFormProps {
  handleLogin: () => void; // handleLogin プロパティの型定義
}

const SignUpForm: React.FC<SignUpFormProps> = ({ handleLogin }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [username,setUsername] = useState('');
  const [userEmail, setUserEmail] = useState('');
  const [cookies, setCookie] = useCookies(['userEmail']);

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
    // ...
  
    try {
      const res = await createUserWithEmailAndPassword(fireAuth, email, password);
      const user = res.user;
      alert("新規登録ユーザー: " + username); // ユーザー名を表示する
      handleLogin(); // ログイン成功時に handleLogin を呼び出す
      setCookie('userEmail', email);

      
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
        <input type="username" value={username} onChange={handleUsernameChange}/>
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
