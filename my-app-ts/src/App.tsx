import React, { useState, ChangeEvent } from 'react';
import './App.css';
import { createUserWithEmailAndPassword, signInWithEmailAndPassword, signOut, onAuthStateChanged, Auth } from "firebase/auth";
import { fireAuth } from "./firebase";

export const LoginForm: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const signUpWithEmail = (): void => {
    // メール/パスワードで新規アカウントを作成する
    createUserWithEmailAndPassword(fireAuth, email, password)
      .then(res => {
        const user = res.user;
        alert("新規アカウントを作成しました。ログインユーザー: " + user?.displayName);
      })
      .catch(err => {
        const errorMessage = err.message;
        alert(errorMessage);
      });
  };

  const signInWithEmail = (): void => {
    // メール/パスワードでログインする
    signInWithEmailAndPassword(fireAuth, email, password)
      .then(res => {
        const user = res.user;
        alert("ログインユーザー: " + user?.displayName);
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
        メール/パスワードで新規登録
      </button>
      <br />
      <button onClick={signInWithEmail}>
        メール/パスワードでログイン
      </button>
      <br />
      <button onClick={signOutWithGoogle}>
        ログアウト
      </button>
    </div>
  );
};

export const Contents = () => {
  return (
    <div>
      <h2>コンテンツ</h2>
      <p>ログインしているユーザーのみが見ることができるコンテンツです。</p>
    </div>
  );
};

const App: React.FC = () => {
  const [loginUser, setLoginUser] = useState<Auth["currentUser"]>(fireAuth.currentUser);

  onAuthStateChanged(fireAuth, user => {
    setLoginUser(user);
  });

  return (
    <>
      <LoginForm />
      {loginUser ? <Contents /> : null}
    </>
  );
};

export default App;
