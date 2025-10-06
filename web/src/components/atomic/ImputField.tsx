const InputField = ({ label, type, register, name, error, ...rest }) => (
  <div style={{ marginBottom: "1rem" }}>
    <label htmlFor={name}>{label}</label>
    <input id={name} type={type} {...register} {...rest} />
    {error && <p style={{ color: "red" }}>{error.message}</p>}
  </div>
);

export default InputField;