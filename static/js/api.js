const request = async (url, options = {}) => {
  const response = await fetch(url, {
    headers: { "Content-Type": "application/json", ...options.headers },
    ...options,
  });
  if (!response.ok) {
    const message = await response.text();
    throw new Error(message || `Error ${response.status}`);
  }
  return response.status === 204 ? null : response.json();
};

const resource = (name) => ({
  list: () => request(`/api/${name}`),
  create: (data) => request(`/api/${name}`, { method: "POST", body: JSON.stringify(data) }),
  remove: (id) => request(`/api/${name}/${encodeURIComponent(id)}`, { method: "DELETE" }),
});

export const api = {
  pacientes: resource("pacientes"),
  enfermeros: resource("enfermeros"),
  familiares: resource("familiares"),
  avisos: resource("avisos"),
};
