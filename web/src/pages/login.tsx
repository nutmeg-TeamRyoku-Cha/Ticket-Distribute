import React, { useState } from 'react';

function RegistrationForm() {
  const [nickname, setNickname] = useState('');
  const [birthdate, setBirthdate] = useState('');
  const [attendees, setAttendees] = useState(1);

  const handleSubmit = (e) => {
    e.preventDefault();
    alert(`登録完了！\nニックネーム: ${nickname}\n生年月日: ${birthdate}\n来場者人数: ${attendees}`);
  };

  return (
    <div style={styles.container}>
      <h1 style={styles.title}>技大祭へようこそ！！</h1>
      <p style={styles.subtitle}>来場者登録をお願いします</p>
      <form onSubmit={handleSubmit} style={styles.form}>
        <label style={styles.label}>
          ニックネーム
          <input
            type="text"
            value={nickname}
            onChange={(e) => setNickname(e.target.value)}
            style={styles.input}
            required
          />
        </label>
        <label style={styles.label}>
          生年月日
          <input
            type="date"
            value={birthdate}
            onChange={(e) => setBirthdate(e.target.value)}
            style={styles.input}
            required
          />
        </label>
        <label style={styles.label}>
          来場者人数
          <input
            type="number"
            value={attendees}
            onChange={(e) => setAttendees(e.target.value)}
            style={styles.input}
            min="1"
            required
          />
        </label>
        <button type="submit" style={styles.button}>登録する</button>
      </form>
    </div>
  );
}

const styles = {
  container: {
    maxWidth: '400px',
    margin: '50px auto',
    padding: '20px',
    border: '1px solid #ccc',
    borderRadius: '10px',
    fontFamily: 'sans-serif',
    backgroundColor: '#f9f9f9',
  },
  title: {
    textAlign: 'center',
    color: '#333',
  },
  subtitle: {
    textAlign: 'center',
    marginBottom: '20px',
    color: '#666',
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
  },
  label: {
    marginBottom: '15px',
    fontWeight: 'bold',
  },
  input: {
    padding: '8px',
    fontSize: '16px',
    marginTop: '5px',
    borderRadius: '5px',
    border: '1px solid #ccc',
  },
  button: {
    padding: '10px',
    fontSize: '16px',
    backgroundColor: '#007bff',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
};

export default RegistrationForm;