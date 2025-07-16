function SeatResult({ seats }) {
  return (
    <div style={{ marginTop: '20px' }}>
      <h3>Generated Seats:</h3>
      <ul>
        {seats.map((seat) => (
          <li key={seat}>{seat}</li>
        ))}
      </ul>
    </div>
  );
}

export default SeatResult;