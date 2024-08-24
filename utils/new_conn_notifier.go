package net_cat

func notifyClients(notification string) {
	messages <- notification
}
