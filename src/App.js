import React, { useState } from 'react';
import './app.css';

function App() {
  const [count, setCount] = useState(0);
  return (
    <div className="container">
      <header className="glass">
        <h1>Anubhav's Website</h1>
        <button onClick={() => setCount(count + 1)}>Clicks {count}</button>
      </header>
      <section className="glass">
        <h2>Hello folks, I am a Software Developer.</h2>
        <p>This website is just for kicks.</p>
      </section>
      <footer className="glass">
        <p>&copy; {new Date().getFullYear()} Anubhav</p>
      </footer>
    </div>
  );
}

export default App;
