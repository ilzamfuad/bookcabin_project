import { useState } from 'react';
import { checkFlight, generateVoucher } from '../api';
import SeatResult from './SeatResult';

function VoucherForm() {
  const [form, setForm] = useState({
    name: '',
    id: '',
    flightNumber: '',
    date: '',
    aircraft: 'ATR',
  });
  const [error, setError] = useState('');
  const [seats, setSeats] = useState([]);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSeats([]);

    try {
      const checkRes = await checkFlight(form.flightNumber, form.date);
      if (checkRes.data.exists) {
        setError('Voucher already generated for this flight and date.');
        return;
      }

      const genRes = await generateVoucher(form);
      setSeats(genRes.data.seats);
    } catch (err) {
      setError(err.response?.data?.message || 'Something went wrong.');
    }
  };

  return (
    <div style={{ maxWidth: '400px', margin: '0 auto' }}>
      <form onSubmit={handleSubmit}>
        <input name="name" value={form.name} onChange={handleChange} placeholder="Crew Name" /><br />
        <input name="id" value={form.id} onChange={handleChange} placeholder="Crew ID" /><br />
        <input name="flightNumber" value={form.flightNumber} onChange={handleChange} placeholder="Flight Number" /><br />
        <input name="date" value={form.date} onChange={handleChange} placeholder="Date (DD-MM-YY)" /><br />
        <select name="aircraft" value={form.aircraft} onChange={handleChange}>
          <option value="ATR">ATR</option>
          <option value="Airbus 320">Airbus 320</option>
          <option value="Boeing 737 Max">Boeing 737 Max</option>
        </select><br />
        <button type="submit">Generate Vouchers</button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>}
      {seats.length > 0 && <SeatResult seats={seats} />}
    </div>
  );
}

export default VoucherForm;