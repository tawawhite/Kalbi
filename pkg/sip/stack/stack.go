package stack

//import "fmt"
import (
	//"fmt"
	"Kalbi/pkg/sip/message"
	"Kalbi/pkg/transport"
	"Kalbi/pkg/sip/transaction"
	"Kalbi/pkg/sip/dialog"
	"Kalbi/pkg/log"
)

//NewSipStack  creates new sip stack
func NewSipStack(Name string) *SipStack {
	stack := new(SipStack)
	stack.Name = Name
	stack.TxMng = transaction.NewTransactionManager()
	
	return stack
}


//SipStack has multiple protocol listning points
type SipStack struct {
	Name              string
	ListeningPoints   []transport.ListeningPoint
	OutputPoint       chan message.SipMsg
	InputPoint        chan message.SipMsg
	Alive             bool
	TxMng             *transaction.TransactionManager
	Dialogs           []dialog.Dialog
	RequestChannel    chan transaction.Transaction
	ResponseChannel   chan transaction.Transaction

}


//CreateListenPoint creates listening point to the event dispatcher
func (ed *SipStack) CreateListenPoint(protocol string, host string, port int) transport.ListeningPoint {
	listenpoint := transport.NewTransportListenPoint(protocol, host, port)
	ed.ListeningPoints = append(ed.ListeningPoints, listenpoint)
    return listenpoint
}

func (ed *SipStack) CreateRequestsChannel() chan transaction.Transaction {
	Channel := make(chan transaction.Transaction)
	ed.RequestChannel =  Channel
	ed.TxMng.RequestChannel = Channel
	return Channel 
}

func (ed *SipStack) CreateResponseChannel() chan transaction.Transaction {
	Channel := make(chan transaction.Transaction)
	ed.ResponseChannel = Channel
	ed.TxMng.ResponseChannel = Channel
	return Channel 
}

func (ed *SipStack) IsAlive() bool {
    return ed.Alive
}

func (ed *SipStack) Stop(){
	log.Log.Info("Stopping SIPStack...")
	ed.Alive = false
}

//Start starts the sip stack
func (ed *SipStack) Start() {
	log.Log.Info("Starting SIPStack...")
	ed.TxMng.ListeningPoint = ed.ListeningPoints[0]
    ed.Alive = true
	for ed.Alive == true {
		for _, listeningPoint := range ed.ListeningPoints {
			msg := listeningPoint.Read()
		    ed.TxMng.Handle(msg)
			
		}
       
	}
}
