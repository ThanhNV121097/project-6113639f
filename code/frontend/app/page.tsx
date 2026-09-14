import { EditablePersistedGreeting } from "../components/EditablePersistedGreeting";
import { readGreeting } from "../lib/editable-persisted-greeting";

export default async function Page() {
  let initialGreeting: string | undefined;

  try {
    initialGreeting = (await readGreeting()).greeting;
  } catch {
    initialGreeting = undefined;
  }

  return <EditablePersistedGreeting initialGreeting={initialGreeting} />;
}
