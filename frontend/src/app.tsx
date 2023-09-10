import { h } from "preact";
import { useState, useEffect } from "preact/hooks";

const COLS = 10;
const ROWS = 20;
const BLOCK_SIZE = 30;

const PIECES = [
  [[1, 1, 1, 1]],
  [
    [1, 1],
    [1, 1],
  ],
  [
    [1, 1, 0],
    [0, 1, 1],
  ],
  [
    [0, 1, 1],
    [1, 1],
  ],
  [
    [1, 1, 1],
    [1, 0, 0],
  ],
  [
    [1, 1, 1],
    [0, 0, 1],
  ],
  [
    [1, 1, 1],
    [0, 1, 0],
  ],
];

function App() {
  const [grid, setGrid] = useState(initGrid());
  const [piece, setPiece] = useState(newPiece());
  const [gameOver, setGameOver] = useState(false);

  function initGrid() {
    let initialGrid: number[] = [];
    for (let r = 0; r < ROWS; r++) {
      initialGrid[r] = [];
      for (let c = 0; c < COLS; c++) {
        initialGrid[r][c] = 0;
      }
    }
    return initialGrid;
  }

  function newPiece() {
    let newP = PIECES[Math.floor(Math.random() * PIECES.length)];
    return {
      shape: newP,
      pos: { x: Math.floor(COLS / 2) - 1, y: 0 },
    };
  }

  useEffect(() => {
    const handleKeyPress = (event) => {
      if (gameOver) return;

      let newPiece = { ...piece };

      switch (event.keyCode) {
        case 37: // 左キー
          newPiece.pos.x -= 1;
          break;
        case 39: // 右キー
          newPiece.pos.x += 1;
          break;
        case 40: // 下キー
          newPiece.pos.y += 1;
          break;
        default:
          return;
      }

      // コリジョンチェックなどをここに追加
      // 例: newPieceがgridの外に出るか、他のブロックと衝突するか

      setPiece(newPiece);
    };

    const gameLoop = setInterval(() => {
      // 自動でブロックを下に移動
      let newPiece = { ...piece };
      newPiece.pos.y += 1;
      // コリジョンチェックなど
      setPiece(newPiece);
    }, 1000);

    document.addEventListener("keydown", handleKeyPress);

    return () => {
      clearInterval(gameLoop);
      document.removeEventListener("keydown", handleKeyPress);
    };
  }, [grid, piece, gameOver]);

  return (
    <div>
      {grid.map((row) => (
        <div>
          {row.map((cell) => (
            <div
              style={{
                width: BLOCK_SIZE,
                height: BLOCK_SIZE,
                border: "1px solid black",
                backgroundColor: cell ? "blue" : "white",
              }}
            />
          ))}
        </div>
      ))}
    </div>
  );
}

export default App;
