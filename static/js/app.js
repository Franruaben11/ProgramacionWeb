import { api } from "./api.js";
import { fillSelect, formData, idOf, nameOf, renderRecords, showStatus, updateCount } from "./dom.js";

const state = { pacientes: [], enfermeros: [], familiares: [], avisos: [] };
const status = document.querySelector("#status");
const list = (type) => document.querySelector(`#${type}-list`);

const noticeLabel = (aviso) => {
  const patient = state.pacientes.find((item) => idOf(item, "paciente") === aviso.id_paciente);
  return `${aviso.nombre} · ${patient ? nameOf(patient) : `Paciente #${aviso.id_paciente}`}`;
};

const render = (type) => {
  renderRecords(list(type), state[type], type, type === "avisos" ? noticeLabel : nameOf);
  updateCount(type, state[type].length);
};

const load = async () => {
  const types = Object.keys(state);
  const values = await Promise.all(types.map((type) => api[type].list()));
  types.forEach((type, index) => {
    state[type] = Array.isArray(values[index]) ? values[index] : [];
    render(type);
  });
  fillSelect(document.querySelector("#aviso-paciente"), state.pacientes, "paciente");
  fillSelect(document.querySelector("#aviso-enfermero"), state.enfermeros, "enfermero");
};

const submit = async (event) => {
  event.preventDefault();
  const form = event.currentTarget;
  const type = form.dataset.form;
  const data = formData(form);
  if (type === "avisos") {
    data.id_paciente = Number(data.id_paciente);
    data.id_enfermero = Number(data.id_enfermero);
    data.fecha_hora_inicio = new Date(data.fecha_hora_inicio).toISOString();
    data.fecha_hora_fin = new Date(data.fecha_hora_fin).toISOString();
  }
  try {
    await api[type].create(data);
    form.reset();
    await load();
    showStatus(status, `${type[0].toUpperCase()}${type.slice(1)} creado correctamente.`);
  } catch (error) {
    showStatus(status, error.message, true);
  }
};

const remove = async (event) => {
  const button = event.target.closest("[data-delete]");
  if (!button) return;
  const type = button.closest(".panel-card").id.replace("-panel", "");
  if (!window.confirm("¿Querés eliminar este registro?")) return;
  try {
    await api[type].remove(button.dataset.delete);
    await load();
    showStatus(status, "Registro eliminado correctamente.");
  } catch (error) {
    showStatus(status, error.message, true);
  }
};

document.querySelectorAll("form[data-form]").forEach((form) => form.addEventListener("submit", submit));
document.querySelectorAll(".record-list").forEach((element) => element.addEventListener("click", remove));
load().catch((error) => showStatus(status, `No se pudieron cargar los registros: ${error.message}`, true));
