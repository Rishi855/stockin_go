import React, { useMemo } from "react";

/**
 * rows: array of objects
 * Renders dynamic header from keys of first row (or union of keys)
 */
export default function DynamicTable({ rows = [] }) {
  const { columns, normalizedRows } = useMemo(() => {
    const colsSet = new Set();
    rows.forEach(row => {
      if (row && typeof row === "object") {
        Object.keys(row).forEach(k => colsSet.add(k));
      }
    });
    const columns = Array.from(colsSet);

    // normalize values to strings for safe rendering
    const normalizedRows = rows.map(row => {
      const out = {};
      columns.forEach(c => {
        let v = row?.[c];
        if (v === null || v === undefined) v = "";
        else if (typeof v === "object") v = JSON.stringify(v);
        else v = String(v);
        out[c] = v;
      });
      return out;
    });

    return { columns, normalizedRows };
  }, [rows]);

  if (!rows || rows.length === 0) {
    return <div className="empty-table">No data to display</div>;
  }

  return (
    <div className="dynamic-table-container" >
      <div className="table-scroll">
        <table className="dynamic-table">
          <thead>
            <tr>
              {columns.map(col => <th key={col}>{col}</th>)}
            </tr>
          </thead>
          <tbody>
            {normalizedRows.map((r, i) => (
              <tr key={i}>
                {columns.map(col => <td key={col + i} title={r[col]}>{r[col]}</td>)}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
