import React, { useEffect, useState } from "react";

const N = 9;
function randomMole() { return Math.floor(Math.random() * N); }

export default function WhacAMole() {
  const [mole, setMole] = useState(randomMole());
  const [score, setScore] = useState(0);
  const [started, setStarted] = useState(false);
  const [timer, setTimer] = useState(30);

  useEffect(() => {
    if (!started) return;
    if (timer <= 0) return;
    const id = setTimeout(() => setTimer(t => t - 1), 1000);
    return () => clearTimeout(id);
  }, [timer, started]);

  useEffect(() => {
    if (!started) return;
    const id = setInterval(() => setMole(randomMole()), 600);
    return () => clearInterval(id);
  }, [started]);

  function whack(i: number) {
    if (i === mole) setScore(s => s + 1);
    setMole(randomMole());
  }

  function start() {
    setScore(0); setTimer(30); setStarted(true);
  }

  return (
    <div style={{ margin: 40 }}>
      <h1>Whac-A-Mole</h1>
      <button onClick={start} disabled={started}>Start Game</button>
      <h2>Time: {timer}s &nbsp; Score: {score}</h2>
      <div style={{ display: "grid", gridTemplateColumns: "repeat(3,70px)", gap: 8 }}>
        {Array.from({ length: N }).map((_, i) =>
          <button key={i} onClick={() => whack(i)}
            style={{
              width: 70, height: 70, fontSize: 30,
              background: mole === i ? "#e33" : "#ccc"
            }}>
            {mole === i ? "🐹" : ""}
          </button>
        )}
      </div>
      {started && timer <= 0 && <h2 style={{ color: "green" }}>Game Over! Final Score: {score}</h2>}
    </div>
  );
}