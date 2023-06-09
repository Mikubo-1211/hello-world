import React, { useState, ChangeEvent } from 'react';
import { signInWithEmailAndPassword, signOut } from "./firebase";
import { Link } from 'react-router-dom';
import { fireAuth } from "./firebase";
import { useCookies } from 'react-cookie';


interface LoginFormProps {
  handleLogin: () => void; // handleLogin プロパティの型定義
}


const LoginForm: React.FC<LoginFormProps> = ({ handleLogin }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [username, setUsername] = useState('');
  const [userEmail, setUserEmail] = useState('');
  const [cookies, setCookie] = useCookies(['userEmail']);


  const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const signInWithEmail = (): void => {
    // メール/パスワードでログインする
    signInWithEmailAndPassword(fireAuth, email, password)
      .then(res => {
        handleLogin(); 
        setCookie('userEmail', res.user?.email || '');
      })
      .catch(err => {
        const errorMessage = err.message;
        alert(errorMessage);
      });
  };

  const signOutWithGoogle = (): void => {
    signOut(fireAuth).then(() => {
      alert("ログアウトしました");
    }).catch(err => {
      alert(err);
    });
  };

  

  return (
    <div>
      <h2>ログイン</h2>
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
      <button onClick={signInWithEmail}>
        ログイン
      </button>
      <br />
      <button onClick={signOutWithGoogle}>
        ログアウト
      </button>
      <br />
      <p>新規登録はこちら: <Link to="/signup">新規登録</Link></p>
    </div>
  );
};
export default LoginForm;