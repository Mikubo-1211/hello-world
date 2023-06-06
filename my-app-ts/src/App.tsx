import React, {useState} from 'react';
import './App.css';
import { signInWithPopup, GoogleAuthProvider, signOut, onAuthStateChanged } from "firebase/auth";
import { fireAuth } from "./firebase";


export const LoginForm: React.FC = () => {
  /**
   * googleでログインする
   */
  const signInWithGoogle = (): void => {
    // Google認証プロバイダを利用する
    const provider = new GoogleAuthProvider();

    // ログイン用のポップアップを表示
    signInWithPopup(fireAuth, provider)
      .then(res => {
        const user = res.user;
        alert("ログインユーザー: " + user.displayName);
      })
      .catch(err => {
        const errorMessage = err.message;
        alert(errorMessage);
      });
  };

  /**
   * ログアウトする
   */
  const signOutWithGoogle = (): void => {
    signOut(fireAuth).then(() => {
      alert("ログアウトしました");
    }).catch(err => {
      alert(err);
    });
  };


  return (
    <div>
      <button onClick={signInWithGoogle}>
        Googleでログイン
      </button>
      <br/>
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




const App = () => {
  // stateとしてログイン状態を管理する。ログインしていないときはnullになる。
  const [loginUser, setLoginUser] = useState(fireAuth.currentUser);
  
  // ログイン状態を監視して、stateをリアルタイムで更新する
  onAuthStateChanged(fireAuth, user => {
    setLoginUser(user);
  });
  
  return (
    <>
      <LoginForm />
      {/* ログインしていないと見られないコンテンツは、loginUserがnullの場合表示しない */}
      {loginUser ? <Contents /> : null}
    </>
  );
};

export default App
