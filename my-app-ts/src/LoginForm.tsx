import React, { useState, ChangeEvent, useContext } from 'react';
import { signInWithEmailAndPassword, signOut } from "./firebase";
import { Link } from 'react-router-dom';
import { fireAuth } from "./firebase";
import { UserContext } from './UserContext';
import './LoginForm.css';

interface LoginFormProps {
  handleLogin: () => void; // handleLogin プロパティの型定義
}

const LoginForm: React.FC<LoginFormProps> = ({ handleLogin }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { setUserEmail } = useContext(UserContext); // UserContextからuserEmailとsetUserEmailを取得
  const [isButtonDisabled, setIsButtonDisabled] = useState(false);
  

  const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const signInWithEmail = (): void => {
    setIsButtonDisabled(true);

    setTimeout(() => {
      setIsButtonDisabled(false);
    }, 2000);
    // メール/パスワードでログインする
    signInWithEmailAndPassword(fireAuth, email, password)
      .then(res => {
        handleLogin(); 
        setUserEmail(res.user?.email || ''); // userEmailを更新
      })
      .catch(err => {
        const errorMessage = err.message;
        alert(errorMessage);
      });
  };

  const signOutWithGoogle = (): void => {
    signOut(fireAuth).then(() => {
      alert("ログアウトしました");
      setUserEmail(''); // userEmailをリセット
    }).catch(err => {
      alert(err);
    });
  };

  return (
    <div className="LoginForm">
      <h1>ログイン</h1>
      <label className="labelText">
        メールアドレス
        <br/>
        <input type="email" value={email} onChange={handleEmailChange} className="inputField"/>
      </label>
      <br />
      <label className="labelText">
        パスワード
        <br/>
        <input type="password" value={password} onChange={handlePasswordChange} className="inputField"/>
      </label>
      <br />
      <button onClick={signInWithEmail} disabled={isButtonDisabled} className="submitButton">
        ログイン
      </button>
      <br />
      <p className="inputField" ><Link to="/signup">新規登録はこちら</Link></p>
    </div>
  );
};

export default LoginForm;
