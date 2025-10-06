import React from 'react';
import Button from '../components/atomic/Button';

const TicketList = () => {
  const handleRegistration = () => {
    alert('登録されました');
  };

  const handleLogin = ()=> {
    alert('ログインしました');
  };

  return (
    <div>
      <Button label="登録する" variant="registration" onClick={handleRegistration} />
      <Button label="ログイン" variant="login" onClick={handleLogin} />
    </div>
  );
};

export default TicketList;