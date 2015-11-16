package config

import (
	"net/url"

	"github.com/twstrike/otr3"

	. "gopkg.in/check.v1"
)

type AccountXmppSuite struct{}

var _ = Suite(&AccountXmppSuite{})

func (s *AccountXmppSuite) Test_Account_Is_recognizesJids(c *C) {
	a := &Account{Account: "hello@bar.com"}
	c.Check(a.Is("foo"), Equals, false)
	c.Check(a.Is("hello@bar.com"), Equals, true)
	c.Check(a.Is("hello@bar.com/foo"), Equals, true)
}

func (s *AccountXmppSuite) Test_Account_ShouldEncryptTo(c *C) {
	a := &Account{Account: "hello@bar.com", AlwaysEncrypt: false, AlwaysEncryptWith: []string{"one@foo.com", "two@foo.com"}}
	a2 := &Account{Account: "hello@bar.com", AlwaysEncrypt: true, AlwaysEncryptWith: []string{"one@foo.com", "two@foo.com"}}
	c.Check(a.ShouldEncryptTo("foo"), Equals, false)
	c.Check(a.ShouldEncryptTo("hello@bar.com"), Equals, false)
	c.Check(a.ShouldEncryptTo("one@foo.com"), Equals, true)
	c.Check(a.ShouldEncryptTo("two@foo.com"), Equals, true)
	c.Check(a.ShouldEncryptTo("two@foo.com/blarg"), Equals, true)
	c.Check(a2.ShouldEncryptTo("foo"), Equals, true)
	c.Check(a2.ShouldEncryptTo("hello@bar.com"), Equals, true)
}

func (s *AccountXmppSuite) Test_EnsureTorProxy_AddsTorProxy(c *C) {
	torAddress := "127.0.0.1:9050"

	a := &Account{
		RequireTor: true,
	}

	a.EnsureTorProxy(torAddress)

	c.Check(len(a.Proxies), Equals, 1)
	proxy, _ := url.Parse(a.Proxies[0])
	c.Check(proxy.Host, Equals, torAddress)
}

func (s *AccountXmppSuite) Test_EnsureTorProxy_AddsTorProxyToTheLastPosition(c *C) {
	torAddress := "127.0.0.1:9050"
	existingTorProxy := "socks5://127.0.0.1:9080"

	a := &Account{
		RequireTor: true,
		Proxies:    []string{existingTorProxy},
	}

	a.EnsureTorProxy(torAddress)

	c.Check(len(a.Proxies), Equals, 2)
	c.Check(a.Proxies[0], Equals, existingTorProxy)

	proxy, _ := url.Parse(a.Proxies[1])
	c.Check(proxy.Host, Equals, torAddress)
}

func (s *AccountXmppSuite) Test_EnsureTorProxy_DoNotOverrideExistingTorConfig(c *C) {
	torAddress := "127.0.0.1:9050"
	existingTorProxy := "socks5://foo:bar@" + torAddress

	a := &Account{
		RequireTor: true,
		Proxies:    []string{existingTorProxy},
	}

	a.EnsureTorProxy(torAddress)
	c.Check(a.Proxies, DeepEquals, []string{existingTorProxy})
}

func (s *AccountXmppSuite) Test_NewAccount_ReturnsNewAccountWithSafeDefaults(c *C) {
	a, err := NewAccount()

	c.Check(err, IsNil)
	c.Check(a.RequireTor, Equals, true)
	c.Check(a.PrivateKey, NotNil)
	c.Check(a.AlwaysEncrypt, Equals, true)
	c.Check(a.OTRAutoStartSession, Equals, true)
	c.Check(a.OTRAutoTearDown, Equals, true)
}

func (s *AccountXmppSuite) Test_SetOTRPoliciesFor_SetupOTRPolicies(c *C) {
	a, _ := NewAccount()
	conv := &otr3.Conversation{}

	expectedConv := &otr3.Conversation{}
	expectedPolicies := expectedConv.Policies
	expectedPolicies.AllowV2()
	expectedPolicies.AllowV3()
	expectedPolicies.SendWhitespaceTag()
	expectedPolicies.WhitespaceStartAKE()
	expectedPolicies.RequireEncryption()
	expectedPolicies.ErrorStartAKE()

	a.SetOTRPoliciesFor("someon@jabber.com", conv)
	c.Check(conv.Policies, Equals, expectedPolicies)
}

func (s *AccountXmppSuite) Test_SetOTRPoliciesFor_SetupOTRPoliciesWithOptionalEncription(c *C) {
	a, _ := NewAccount()
	a.AlwaysEncrypt = false
	conv := &otr3.Conversation{}

	expectedConv := &otr3.Conversation{}
	expectedPolicies := expectedConv.Policies
	expectedPolicies.AllowV2()
	expectedPolicies.AllowV3()
	expectedPolicies.SendWhitespaceTag()
	expectedPolicies.WhitespaceStartAKE()

	a.SetOTRPoliciesFor("someon@jabber.com", conv)
	c.Check(conv.Policies, Equals, expectedPolicies)
}

func (s *AccountXmppSuite) Test_EnsurePrivateKey_DoesNotUpdateIfKeyExists(c *C) {
	a, _ := NewAccount()
	changed, err := a.EnsurePrivateKey()

	c.Check(err, IsNil)
	c.Check(changed, Equals, false)
}

func (s *AccountXmppSuite) Test_EnsurePrivateKey_GeneratePrivateKeyIfMissing(c *C) {
	a := &Account{}
	changed, err := a.EnsurePrivateKey()

	c.Check(err, IsNil)
	c.Check(changed, Equals, true)
	c.Check(a.PrivateKey, NotNil)
}

func (s *AccountXmppSuite) Test_ID_generatesID(c *C) {
	a := &Account{}
	c.Check(a.ID(), Not(HasLen), 0)
}

func (s *AccountXmppSuite) Test_ID_doesNotChangeID(c *C) {
	a := &Account{
		id: "existing",
	}
	c.Check(a.ID(), Equals, "existing")
}
