import { useState } from 'react'
import logo from './logo.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      <header className="App-header">
        <title>home</title>
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          <button type="button" onClick={() => setCount(count + 1)}>
           clicks: {count}
          </button>
          <button type="button" onClick={()=> setCount(0)}>reset</button>
        </p>
        <p>
          Edit <code>App.tsx</code> and save to test HMR updates.
        </p>
      </header>
    </div>
  )
}

export default App
