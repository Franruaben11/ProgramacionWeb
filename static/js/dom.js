const escapeHtml = (value) => String(value ?? "").replace(/[&<>"']/g, (char) => ({
  "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#039;",
}[char]));

const idKeys = {
  pacientes: "id_paciente",
  enfermeros: "id_enfermero",
  familiares: "id_familiar",
  avisos: "id_aviso",
  paciente: "id_paciente",
  enfermero: "id_enfermero",
  familiar: "id_familiar",
  aviso: "id_aviso",
};

export const idOf = (record, type) => record[idKeys[type] || `id_${type}`] ?? record.id;
export const nameOf = (record) => record.nombre || `Registro #${record.id}`;
export const formData = (form) => Object.fromEntries(new FormData(form).entries());

export const renderRecords = (list, records, type, label = nameOf) => {
  list.replaceChildren();
  records.forEach((record) => {
    const id = idOf(record, type);
    const item = document.createElement("li");
    item.innerHTML = `<span><strong>#${escapeHtml(id)}</strong> ${escapeHtml(label(record))}</span>
      <button class="delete-button" type="button" data-delete="${escapeHtml(id)}" aria-label="Eliminar ${escapeHtml(label(record))}">Eliminar</button>`;
    list.append(item);
  });
};

export const updateCount = (type, count) => {
  const element = document.querySelector(`#${type}-count`);
  if (element) element.textContent = count;
};

export const fillSelect = (select, records, type) => {
  const current = select.value;
  select.replaceChildren(new Option(`Seleccioná un ${type === "paciente" ? "paciente" : "enfermero"}`, ""));
  records.forEach((record) => select.add(new Option(`#${idOf(record, type)} · ${nameOf(record)}`, idOf(record, type))));
  if ([...select.options].some((option) => option.value === current)) select.value = current;
};

export const showStatus = (element, message, error = false) => {
  element.textContent = message;
  element.className = `status${error ? " status-error" : " status-success"}`;
};
