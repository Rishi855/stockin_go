import React, { useState } from "react";
import DynamicTable from "./components/DynamicTable";
import Toast from "./components/Toast";


export default function App() {
  const [sql, setSql] = useState("");
  const [rows, setRows] = useState([]);
  const [offset, setOffset] = useState(0);
  const [pageSize] = useState(10); // change if you want other page size
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [lastStockUpdate, setLastStockUpdate] = useState(null);
  const [lastGrowwUpdate, setLastGrowwUpdate] = useState(null);
  const [lastNewsUpdate, setLastNewsUpdate] = useState(null);

  // per-button loading flags
  const [updatingStock, setUpdatingStock] = useState(false);
  const [updatingGroww, setUpdatingGroww] = useState(false);
  const [updatingNews, setUpdatingNews] = useState(false);

  const [toastMsg, setToastMsg] = useState("");


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

  async function triggerUpdate(url, setTimestamp, label, loaderSetter) {
    try {
      loaderSetter(true);               // disable button while running
      // perform request (we don't care about response body)
      const resp = await fetch(url, { method: "POST", headers: { "Content-Type": "application/json" } });

      if (!resp.ok) {
        const txt = await resp.text().catch(() => "");
        const errMsg = `${label} update failed (${resp.status}) ${txt}`;
        setToastMsg(errMsg);
        return;
      }

      // success: set timestamp and show toast AFTER response
      setTimestamp(new Date().toLocaleString());
      setToastMsg(`${label} update started (server accepted).`);
    } catch (err) {
      console.error(err);
      setToastMsg(`Failed: ${label} update`);
    } finally {
      loaderSetter(false);              // re-enable button
    }
  }

  const updateStockDb = () =>
    triggerUpdate(
      "http://localhost:8080/api/database/scrap/stock",
      setLastStockUpdate,
      "Stock database",
      setUpdatingStock
    );

  const updateGrowwDb = () =>
    triggerUpdate(
      "http://localhost:8080/api/database/scrap/groww",
      setLastGrowwUpdate,
      "Groww database",
      setUpdatingGroww
    );

  const updateNewsDb = () =>
    triggerUpdate(
      "http://localhost:8080/api/database/scrap/news",
      setLastNewsUpdate,
      "News database",
      setUpdatingNews
    );

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
      <div className="update-buttons">
        <div className="update-row-inline">
          <button
            onClick={updateGrowwDb}
            className="update-btn"
            disabled={updatingGroww}
            title="Starts large groww DB update"
          >
            {updatingGroww ? "Updating groww..." : "Scrap Groww Database(main)"}
          </button>
          <span className="timestamp">{lastGrowwUpdate ? `Last: ${lastGrowwUpdate}` : ""}</span>

          <button
            onClick={updateStockDb}
            className="update-btn"
            disabled={updatingStock}
            title="Starts large stock DB update"
          >
            {updatingStock ? "Updating stock..." : "Update Stock Database"}
          </button>
          <span className="timestamp">{lastStockUpdate ? `Last: ${lastStockUpdate}` : ""}</span>


          <button
            onClick={updateNewsDb}
            className="update-btn"
            disabled={updatingNews}
            title="Starts large news DB update"
          >
            {updatingNews ? "Updating news..." : "Update News Database"}
          </button>
          <span className="timestamp">{lastNewsUpdate ? `Last: ${lastNewsUpdate}` : ""}</span>
        </div>
      </div>

      <Toast message={toastMsg} onClose={() => setToastMsg("")} />


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
          <button
            className="page-btn"
            onClick={onPrev}
            disabled={loading || offset === 0}
          >
            ⬅ Previous
          </button>

          <button
            className="page-btn"
            onClick={onNext}
            disabled={loading || rows.length < pageSize}
            title="Next page"
          >
            Next ➡
          </button>
        </div>

        {error && <div className="error">Error: {error}</div>}

        <div className="table-wrapper">
          <DynamicTable rows={rows} />
        </div>

      {/* <div className="table-bottom-controls">
        <button
          className="page-btn"
          onClick={onPrev}
          disabled={loading || offset === 0}
        >
          ⬅ Previous
        </button>

        <button
          className="page-btn"
          onClick={onNext}
          disabled={loading || rows.length < pageSize}
        >
          Next ➡
        </button>
      </div> */}

      </main>
    </div>
  );
}
