import React, { useState } from "react";
import DynamicTable from "./components/DynamicTable";

export default function App() {
  const [sql, setSql] = useState("");
  const [rows, setRows] = useState([]);
  const [offset, setOffset] = useState(0);
  const [pageSize] = useState(10); // change if you want other page size
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  async function fetchData(currentSql, currentOffset) {
    setLoading(true);
    setError(null);
    try {
      const resp = await fetch("http://localhost:8080/api/query/select", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ sql: currentSql, offset: currentOffset, limit: pageSize }),
      });
      if (!resp.ok) {
        const text = await resp.text();
        throw new Error(`API error: ${resp.status} ${text}`);
      }
      const data = await resp.json();
      // Expecting the API to return an array of rows, e.g. { rows: [...] } or just [...]
      const resultRows = Array.isArray(data) ? data : data.rows ?? data.data ?? [];
      setRows(resultRows);
    } catch (e) {
      console.error(e);
      setError(e.message || "Unknown error");
      setRows([]);
    } finally {
      setLoading(false);
    }
  }

  const onExecute = async () => {
    setOffset(0);
    await fetchData(sql, 0);
  };

  const onNext = async () => {
    const newOffset = offset + pageSize;
    setOffset(newOffset);
    await fetchData(sql, newOffset);
  };

  const onPrev = async () => {
    const newOffset = Math.max(0, offset - pageSize);
    setOffset(newOffset);
    await fetchData(sql, newOffset);
  };

  return (
    <div className="app-root">
      <header className="top-area">
        <div className="top-inner">
          <label htmlFor="sqlBox" className="sr-only">SQL</label>
          <textarea
            id="sqlBox"
            className="sql-textarea"
            placeholder="Write your SQL here..."
            value={sql}
            onChange={(e) => setSql(e.target.value)}
          />
          <div className="controls">
            <button className="execute-btn" onClick={onExecute} disabled={loading || !sql.trim()}>
              {loading ? "Executing..." : "Execute"}
            </button>
            <div className="meta">
              <span>Offset: {offset}</span>
              <span>Page size: {pageSize}</span>
            </div>
          </div>
        </div>
      </header>

      <main className="bottom-area">
        <div className="table-controls">
          <button onClick={onPrev} disabled={loading || offset === 0}>Previous</button>
          <button
            onClick={onNext}
            disabled={loading || rows.length < pageSize}
            title="Disabled when fewer rows returned than page size"
          >
            Next
          </button>
        </div>

        {error && <div className="error">Error: {error}</div>}

        <div className="table-wrapper">
          <DynamicTable rows={rows} />
        </div>

        <div className="table-bottom-controls">
          <button onClick={onPrev} disabled={loading || offset === 0}>Previous</button>
          <button onClick={onNext} disabled={loading || rows.length < pageSize}>Next</button>
        </div>
      </main>
    </div>
  );
}
