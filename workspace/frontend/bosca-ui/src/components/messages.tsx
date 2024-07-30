import {redirect} from "next/navigation";
import {ErrorAlert, InfoAlert, SuccessAlert} from "@/components/alert";

interface MessagesProps {
  messages: Array<Message>
}

interface Message {
  id: number
  type: string
  text: string
}

export function Messages({messages}: MessagesProps) {
  const alerts = []
  for (const message of messages) {
    if (message.type === 'success') {
      alerts.push(
        <SuccessAlert key={"success-" + message.id} title={"Success"} body={message.text}/>
      )
    } else if (message.type === 'info') {
      alerts.push(
        <InfoAlert key={"info-" + message.id} title={"Information"} body={message.text}/>
      )
    } else {
      alerts.push(
        <ErrorAlert key={"error-" + message.id} title={"Error"} body={message.text}/>
      )
    }
  }
  return alerts
}